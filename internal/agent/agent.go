// Package agent holds the baseline and the staged solution reviewers. The loop
// is written by hand: no orchestration framework. Each stage turns a contract
// into a list of reported findings that internal/scoring then grades.
package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/EndPx/saksama/internal/contract"
	"github.com/EndPx/saksama/internal/llm"
	"github.com/EndPx/saksama/internal/scoring"
	"github.com/EndPx/saksama/internal/statutes"
)

// Agent runs reviews against a fixed legal corpus using an LLM client.
type Agent struct {
	Client *llm.Client
	Corpus *statutes.Corpus
}

// New builds an Agent.
func New(c *llm.Client, corpus *statutes.Corpus) *Agent {
	return &Agent{Client: c, Corpus: corpus}
}

// catalog renders the legal corpus as a compact reference for prompts.
func (a *Agent) catalog() string {
	var b strings.Builder
	for _, p := range a.Corpus.Provisions {
		fmt.Fprintf(&b, "- %s | %s %s | tier=%s | deteksi=%s | %s: %s\n",
			p.ID, p.DasarHukum, p.Pasal, p.Tier, p.Deteksi, p.Judul, p.Ringkasan)
	}
	return b.String()
}

// validIDs returns the set of allowed statute ids.
func (a *Agent) validIDs() map[string]bool {
	m := make(map[string]bool, len(a.Corpus.Provisions))
	for _, p := range a.Corpus.Provisions {
		m[p.ID] = true
	}
	return m
}

// extractJSONArray returns the outermost [...] block in s, or "" if none.
func extractJSONArray(s string) string {
	start := strings.Index(s, "[")
	end := strings.LastIndex(s, "]")
	if start < 0 || end < 0 || end < start {
		return ""
	}
	return s[start : end+1]
}

// rawFinding is the JSON shape the model is asked to emit.
type rawFinding struct {
	StatuteID string `json:"statute_id"`
	Section   string `json:"section"`
	Tier      string `json:"tier"`
	Deteksi   string `json:"deteksi"`
	Kutipan   string `json:"kutipan"`
	Deskripsi string `json:"deskripsi"`
}

// parseFindings decodes a JSON array of findings and keeps only those whose
// statute_id is valid. Unparseable input yields an error.
func (a *Agent) parseFindings(text string) ([]scoring.Finding, error) {
	arr := extractJSONArray(text)
	if arr == "" {
		return nil, fmt.Errorf("no JSON array found in model output")
	}
	var raw []rawFinding
	if err := json.Unmarshal([]byte(arr), &raw); err != nil {
		return nil, fmt.Errorf("decode findings: %w", err)
	}
	valid := a.validIDs()
	var out []scoring.Finding
	for _, r := range raw {
		if !valid[r.StatuteID] {
			continue // constrain to the fourteen valid ids
		}
		out = append(out, scoring.Finding{
			StatuteID: r.StatuteID,
			Section:   r.Section,
			Tier:      statutes.Tier(r.Tier),
			Deteksi:   statutes.Detection(r.Deteksi),
			Kutipan:   r.Kutipan,
			Deskripsi: r.Deskripsi,
		})
	}
	return out, nil
}

const jsonSchemaHint = `Balas HANYA dengan array JSON valid, tanpa teks lain, dengan bentuk:
[{"statute_id":"<salah satu id yang valid>","section":"Pasal N atau ABSENT","tier":"<tier>","deteksi":"<ada_klausa|tidak_ada_klausa|konteks>","kutipan":"kutipan harfiah dari kontrak, maks 200 karakter, kosong bila deteksi tidak_ada_klausa","deskripsi":"penjelasan singkat"}]
Jika tidak ada temuan, balas dengan array kosong: []`

// Baseline runs the single free-form review call (no statutes, no schema).
func (a *Agent) Baseline(ctx context.Context, contractText string) (string, llm.Usage, error) {
	resp, err := a.Client.Complete(ctx, llm.Request{
		System: "Anda adalah asisten yang meninjau kontrak kerja di Indonesia.",
		Messages: []llm.Message{{
			Role:    "user",
			Content: "Tinjau kontrak kerja berikut dan sebutkan risiko-risiko yang Anda temukan.\n\n" + contractText,
		}},
		MaxTokens:   6000,
		Temperature: 0,
	})
	return resp.Text, resp.Usage, err
}

// Normalize converts free-form baseline text into structured findings. This
// call is tracked separately and never counted as part of the solution cost.
func (a *Agent) Normalize(ctx context.Context, freeform string) ([]scoring.Finding, llm.Usage, error) {
	sys := "Ubah teks tinjauan bebas menjadi array JSON temuan. Gunakan hanya id statute valid berikut:\n" + a.catalog() + "\n" + jsonSchemaHint
	resp, err := a.Client.Complete(ctx, llm.Request{
		System:      sys,
		Messages:    []llm.Message{{Role: "user", Content: freeform}},
		MaxTokens:   6000,
		Temperature: 0,
	})
	if err != nil {
		return nil, resp.Usage, err
	}
	f, perr := a.parseFindings(resp.Text)
	return f, resp.Usage, perr
}

// Structured (S1) reviews the whole document and returns findings directly as
// constrained JSON.
func (a *Agent) Structured(ctx context.Context, contractText string) ([]scoring.Finding, llm.Usage, error) {
	sys := "Anda meninjau kontrak kerja Indonesia terhadap daftar ketentuan hukum berikut. Laporkan hanya pelanggaran yang relevan.\n" + a.catalog() + "\n" + jsonSchemaHint
	resp, err := a.Client.Complete(ctx, llm.Request{
		System:      sys,
		Messages:    []llm.Message{{Role: "user", Content: contractText}},
		MaxTokens:   6000,
		Temperature: 0,
	})
	if err != nil {
		return nil, resp.Usage, err
	}
	f, perr := a.parseFindings(resp.Text)
	return f, resp.Usage, perr
}

// StructuredPerSection (S2) sends the contract section by section, so long
// documents do not overflow effective context, then merges the findings.
func (a *Agent) StructuredPerSection(ctx context.Context, contractText string) ([]scoring.Finding, llm.Usage, error) {
	preamble, sections := contract.Parse(contractText)
	var all []scoring.Finding
	var usage llm.Usage
	sys := "Anda meninjau SATU pasal dari kontrak kerja Indonesia terhadap daftar ketentuan berikut. Konteks identitas para pihak diberikan. Laporkan hanya pelanggaran pada pasal ini.\n" + a.catalog() + "\n" + jsonSchemaHint
	for _, s := range sections {
		msg := "Identitas para pihak (konteks):\n" + preamble + "\n\n" + s.Heading + "\n" + s.Body
		resp, err := a.Client.Complete(ctx, llm.Request{
			System:      sys,
			Messages:    []llm.Message{{Role: "user", Content: msg}},
			MaxTokens:   6000,
			Temperature: 0,
		})
		usage.InputTokens += resp.Usage.InputTokens
		usage.OutputTokens += resp.Usage.OutputTokens
		usage.CostUSD += resp.Usage.CostUSD
		if err != nil {
			return all, usage, err
		}
		f, perr := a.parseFindings(resp.Text)
		if perr != nil {
			continue // a single unparseable section should not kill the run
		}
		for i := range f {
			if f[i].Section == "" {
				f[i].Section = s.Label() // use the actual section when the model omits it
			}
		}
		all = append(all, f...)
	}
	return all, usage, nil
}
