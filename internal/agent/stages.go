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

// TrajStep is one recorded interaction in a stage trajectory.
type TrajStep struct {
	Label    string
	Prompt   string
	Response string
	Outcome  string
}

// Rejected is a finding dropped by the citation gate, with the reason.
type Rejected struct {
	Finding scoring.Finding
	Reason  string
}

func addUsage(dst *llm.Usage, u llm.Usage) {
	dst.InputTokens += u.InputTokens
	dst.OutputTokens += u.OutputTokens
	dst.CostUSD += u.CostUSD
}

// pp3513Items are the nine mandatory items of PP 35/2021 Pasal 13.
var pp3513Items = []string{
	"nama, alamat perusahaan, dan jenis usaha",
	"nama, jenis kelamin, umur, dan alamat pekerja",
	"jabatan atau jenis pekerjaan",
	"tempat pekerjaan",
	"besaran dan cara pembayaran upah",
	"hak dan kewajiban pengusaha dan pekerja",
	"mulai dan jangka waktu berlakunya PKWT",
	"tempat dan tanggal PKWT dibuat",
	"tanda tangan para pihak",
}

// Checklist (S3) runs one targeted presence check per ada_klausa/konteks
// provision. The agent no longer chooses freely what to look at.
func (a *Agent) Checklist(ctx context.Context, contractText string) ([]scoring.Finding, llm.Usage, []TrajStep, error) {
	var findings []scoring.Finding
	var usage llm.Usage
	var traj []TrajStep
	isPKWTT := contractIsPKWTT(contractText)
	for _, p := range a.Corpus.Provisions {
		if p.Deteksi != statutes.DeteksiAdaKlausa && p.Deteksi != statutes.DeteksiKonteks {
			continue
		}
		if isPKWTT && pkwtOnly[p.ID] {
			continue // provision does not apply to a permanent (PKWTT) contract
		}
		prompt := fmt.Sprintf(
			"Ketentuan hukum:\n%s %s — %s: %s\n\nApakah kontrak berikut memuat klausa yang MELANGGAR ketentuan di atas? "+
				"Jawab melanggar:true HANYA bila ada klausa spesifik dalam kontrak yang jelas melanggar ketentuan ini. "+
				"Bila kontrak patuh, tidak menyinggung hal ini, atau ambigu, jawab melanggar:false. "+
				"Balas HANYA JSON: {\"melanggar\":true|false,\"section\":\"Pasal N (nomor pasal DALAM KONTRAK, bukan pasal peraturan)\",\"kutipan\":\"kutipan harfiah maks 200 karakter\",\"deskripsi\":\"penjelasan singkat\"}.\n\nKONTRAK:\n%s",
			p.DasarHukum, p.Pasal, p.Judul, p.Ringkasan, contractText)
		resp, err := a.Client.Complete(ctx, llm.Request{
			System:      "Anda pemeriksa kepatuhan kontrak kerja Indonesia. Jawab presisi, hanya JSON.",
			Messages:    []llm.Message{{Role: "user", Content: prompt}},
			MaxTokens:   6000,
			Temperature: 0,
		})
		addUsage(&usage, resp.Usage)
		if err != nil {
			traj = append(traj, TrajStep{Label: "checklist:" + p.ID, Prompt: prompt, Outcome: "ERROR (skipped): " + err.Error()})
			continue // a single failed call must not kill the stage
		}
		f, outcome := a.parseSingle(resp.Text, p)
		traj = append(traj, TrajStep{Label: "checklist:" + p.ID, Prompt: prompt, Response: resp.Text, Outcome: outcome})
		if f != nil {
			findings = append(findings, *f)
		}
	}
	return findings, usage, traj, nil
}

// Absence (S4) asks, for each tidak_ada_klausa provision, whether the contract
// satisfies it. PP35-13 is expanded into nine sub-checks. PP35-14 (online
// registration) is intentionally excluded: registration is an administrative
// act, not a contract clause, so its absence from the text is normal and would
// otherwise be a universal false positive.
func (a *Agent) Absence(ctx context.Context, contractText string) ([]scoring.Finding, llm.Usage, []TrajStep, error) {
	var findings []scoring.Finding
	var usage llm.Usage
	var traj []TrajStep
	isPKWTT := contractIsPKWTT(contractText)
	for _, p := range a.Corpus.Provisions {
		if p.Deteksi != statutes.DeteksiTidakAdaKlausa {
			continue
		}
		if p.ID == "PP35-14" {
			continue // administrative act, not a contract clause
		}
		if isPKWTT && pkwtOnly[p.ID] {
			continue // provision does not apply to a permanent (PKWTT) contract
		}
		if p.ID == "PP35-13" {
			f, u, step := a.checkPasal13(ctx, contractText)
			addUsage(&usage, u)
			traj = append(traj, step)
			findings = append(findings, f...)
			continue
		}
		if p.ID == "PP35-27-5" {
			f, u, step := a.checkOvertimeExemption(ctx, contractText)
			addUsage(&usage, u)
			traj = append(traj, step)
			if f != nil {
				findings = append(findings, *f)
			}
			continue
		}
		prompt := fmt.Sprintf(
			"Ketentuan hukum:\n%s %s — %s: %s\n\nApakah kontrak berikut MEMENUHI ketentuan di atas? "+
				"Klausa yang membuat kewajiban menjadi bersyarat atau menghapusnya dihitung sebagai TIDAK memenuhi. "+
				"Balas HANYA JSON: {\"memenuhi\":true|false,\"section\":\"Pasal N atau ABSENT\",\"deskripsi\":\"penjelasan singkat\"}.\n\nKONTRAK:\n%s",
			p.DasarHukum, p.Pasal, p.Judul, p.Ringkasan, contractText)
		resp, err := a.Client.Complete(ctx, llm.Request{
			System:      "Anda pemeriksa kepatuhan kontrak kerja Indonesia. Fokus pada apa yang HILANG. Jawab hanya JSON.",
			Messages:    []llm.Message{{Role: "user", Content: prompt}},
			MaxTokens:   6000,
			Temperature: 0,
		})
		addUsage(&usage, resp.Usage)
		if err != nil {
			traj = append(traj, TrajStep{Label: "absence:" + p.ID, Prompt: prompt, Outcome: "ERROR (skipped): " + err.Error()})
			continue // a single failed call must not kill the stage
		}
		var r struct {
			Memenuhi  bool   `json:"memenuhi"`
			Section   string `json:"section"`
			Deskripsi string `json:"deskripsi"`
		}
		outcome := "memenuhi / tidak ada temuan"
		if err := json.Unmarshal([]byte(extractJSONObject(resp.Text)), &r); err == nil && !r.Memenuhi {
			findings = append(findings, scoring.Finding{
				StatuteID: p.ID, Section: "ABSENT", Tier: p.Tier, Deteksi: p.Deteksi, Deskripsi: r.Deskripsi,
			})
			outcome = "TEMUAN: ketentuan tidak dipenuhi"
		}
		traj = append(traj, TrajStep{Label: "absence:" + p.ID, Prompt: prompt, Response: resp.Text, Outcome: outcome})
	}
	return findings, usage, traj, nil
}

// checkPasal13 runs the nine mandatory-item sub-checks in one call and emits one
// PP35-13 finding per missing item.
func (a *Agent) checkPasal13(ctx context.Context, contractText string) ([]scoring.Finding, llm.Usage, TrajStep) {
	var items strings.Builder
	for i, it := range pp3513Items {
		fmt.Fprintf(&items, "%d. %s\n", i+1, it)
	}
	prompt := fmt.Sprintf(
		"PP 35/2021 Pasal 13 mewajibkan PKWT memuat sembilan hal berikut:\n%s\n"+
			"Untuk setiap nomor, tentukan apakah hal itu ADA dalam kontrak. "+
			"Balas HANYA JSON: {\"missing\":[daftar NOMOR item yang TIDAK ada]}.\n\nKONTRAK:\n%s",
		items.String(), contractText)
	resp, _ := a.Client.Complete(ctx, llm.Request{
		System:      "Anda pemeriksa kelengkapan PKWT. Jawab hanya JSON.",
		Messages:    []llm.Message{{Role: "user", Content: prompt}},
		MaxTokens:   6000,
		Temperature: 0,
	})
	var r struct {
		Missing []int `json:"missing"`
	}
	_ = json.Unmarshal([]byte(extractJSONObject(resp.Text)), &r)
	var findings []scoring.Finding
	for _, n := range r.Missing {
		if n < 1 || n > len(pp3513Items) {
			continue
		}
		findings = append(findings, scoring.Finding{
			StatuteID: "PP35-13", Section: "ABSENT", Tier: statutes.TierMelanggarTanpaSanksi,
			Deteksi: statutes.DeteksiTidakAdaKlausa, Deskripsi: "Item wajib tidak ada: " + pp3513Items[n-1],
		})
	}
	outcome := fmt.Sprintf("%d dari 9 item wajib tidak ada", len(findings))
	return findings, resp.Usage, TrajStep{Label: "absence:PP35-13", Prompt: prompt, Response: resp.Text, Outcome: outcome}
}

// checkOvertimeExemption handles PP35-27-5, which is a violation ONLY when the
// contract actually waives overtime pay without defining the exempt senior job
// category. A contract that never mentions an overtime exemption is compliant,
// so the generic absence check (which flags every contract) is wrong here.
func (a *Agent) checkOvertimeExemption(ctx context.Context, contractText string) (*scoring.Finding, llm.Usage, TrajStep) {
	prompt := "Periksa ketentuan lembur dalam kontrak berikut. " +
		"Apakah kontrak MEMUAT klausa yang menyatakan suatu jabatan dikecualikan atau TIDAK BERHAK atas upah lembur? " +
		"Jika kontrak TIDAK memuat pernyataan pengecualian lembur seperti itu, jawab {\"melanggar\":false}. " +
		"Jika kontrak MEMUAT pernyataan bahwa jabatan tidak berhak lembur TETAPI tidak mendefinisikan golongan jabatan tertentu " +
		"(pemikir, perencana, pelaksana, atau pengendali jalannya perusahaan), jawab {\"melanggar\":true,\"section\":\"Pasal N\",\"kutipan\":\"kutipan harfiah\"}." +
		"\n\nKONTRAK:\n" + contractText
	resp, _ := a.Client.Complete(ctx, llm.Request{
		System:      "Anda pemeriksa ketentuan lembur PKWT. Jawab hanya JSON.",
		Messages:    []llm.Message{{Role: "user", Content: prompt}},
		MaxTokens:   12000,
		Temperature: 0,
	})
	var r struct {
		Melanggar bool   `json:"melanggar"`
		Section   string `json:"section"`
		Kutipan   string `json:"kutipan"`
	}
	_ = json.Unmarshal([]byte(extractJSONObject(resp.Text)), &r)
	if !r.Melanggar {
		return nil, resp.Usage, TrajStep{Label: "absence:PP35-27-5", Prompt: prompt, Response: resp.Text, Outcome: "tidak ada pengecualian lembur / tidak ada temuan"}
	}
	section := r.Section
	if section == "" {
		section = "ABSENT"
	}
	return &scoring.Finding{
		StatuteID: "PP35-27-5", Section: section, Tier: statutes.TierBatalDemiHukum,
		Deteksi: statutes.DeteksiTidakAdaKlausa, Kutipan: r.Kutipan,
		Deskripsi: "Klausa mengecualikan lembur tanpa mendefinisikan golongan jabatan tertentu.",
	}, resp.Usage, TrajStep{Label: "absence:PP35-27-5", Prompt: prompt, Response: resp.Text,
		Outcome: "TEMUAN: pengecualian lembur tanpa definisi golongan jabatan"}
}

// CitationGate (S5) drops any ada_klausa finding whose kutipan is not found
// verbatim (whitespace-normalised) in the contract text.
func (a *Agent) CitationGate(contractText string, findings []scoring.Finding) (kept []scoring.Finding, rejected []Rejected) {
	for _, f := range findings {
		if f.Deteksi != statutes.DeteksiAdaKlausa {
			kept = append(kept, f) // absence/context findings are not gated on quotes
			continue
		}
		if f.Kutipan == "" {
			rejected = append(rejected, Rejected{Finding: f, Reason: "tidak ada kutipan"})
			continue
		}
		if !contract.ContainsQuote(contractText, f.Kutipan) {
			rejected = append(rejected, Rejected{Finding: f, Reason: "kutipan tidak ditemukan dalam kontrak"})
			continue
		}
		kept = append(kept, f)
	}
	return kept, rejected
}

// parseSingle decodes one {"melanggar":...} object into a finding for provision p.
func (a *Agent) parseSingle(text string, p statutes.Provision) (*scoring.Finding, string) {
	var r struct {
		Melanggar bool   `json:"melanggar"`
		Section   string `json:"section"`
		Kutipan   string `json:"kutipan"`
		Deskripsi string `json:"deskripsi"`
	}
	if err := json.Unmarshal([]byte(extractJSONObject(text)), &r); err != nil {
		return nil, "tidak dapat diparse / tidak ada temuan"
	}
	if !r.Melanggar {
		return nil, "tidak melanggar"
	}
	section := r.Section
	if section == "" {
		section = "ABSENT"
	}
	return &scoring.Finding{
		StatuteID: p.ID, Section: section, Tier: p.Tier, Deteksi: p.Deteksi,
		Kutipan: r.Kutipan, Deskripsi: r.Deskripsi,
	}, "TEMUAN: melanggar"
}

// extractJSONObject returns the first balanced, sanitized {...} block, or "{}".
func extractJSONObject(s string) string {
	if o := balancedJSON(s, '{', '}'); o != "" {
		return sanitizeJSON(o)
	}
	return "{}"
}

// RenderTrajectory produces the markdown trajectory for a stage run on one
// contract: every prompt, raw response, kept findings, and rejected findings.
func RenderTrajectory(stage, contractID string, steps []TrajStep, kept []scoring.Finding, rejected []Rejected) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# Trajectory — %s — %s\n\n", stage, contractID)
	for _, s := range steps {
		fmt.Fprintf(&b, "## %s\n\n**Outcome:** %s\n\n", s.Label, s.Outcome)
		b.WriteString("_Prompt:_\n\n```\n" + strings.TrimSpace(s.Prompt) + "\n```\n\n")
		b.WriteString("_Response:_\n\n```\n" + strings.TrimSpace(s.Response) + "\n```\n\n")
	}
	fmt.Fprintf(&b, "## Findings kept (%d)\n\n", len(kept))
	for _, f := range kept {
		fmt.Fprintf(&b, "- %s @ %s\n", f.StatuteID, f.Section)
	}
	fmt.Fprintf(&b, "\n## Findings rejected by citation gate (%d)\n\n", len(rejected))
	for _, r := range rejected {
		fmt.Fprintf(&b, "- %s @ %s — %s\n", r.Finding.StatuteID, r.Finding.Section, r.Reason)
	}
	return b.String()
}
