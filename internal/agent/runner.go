package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/EndPx/saksama/internal/llm"
	"github.com/EndPx/saksama/internal/scoring"
	"github.com/EndPx/saksama/internal/statutes"
)

// Contract is one loaded contract from the corpus.
type Contract struct {
	ID   string
	Text string
}

// LoadContracts reads every data/contracts/cNN/contract.md under dir, sorted by id.
func LoadContracts(dir string) ([]Contract, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var ids []string
	for _, e := range entries {
		if e.IsDir() && strings.HasPrefix(e.Name(), "c") {
			ids = append(ids, e.Name())
		}
	}
	sort.Strings(ids)
	var out []Contract
	for _, id := range ids {
		b, err := os.ReadFile(filepath.Join(dir, id, "contract.md"))
		if err != nil {
			return nil, fmt.Errorf("read %s: %w", id, err)
		}
		out = append(out, Contract{ID: id, Text: string(b)})
	}
	return out, nil
}

func durMs(start time.Time) int64 { return time.Since(start).Milliseconds() }

func toScoringUsage(u llm.Usage) scoring.Usage {
	return scoring.Usage{InputTokens: u.InputTokens, OutputTokens: u.OutputTokens, CostUSD: u.CostUSD}
}

// RunBaseline (S0) runs the free-form review plus the separate normalisation
// call. Normalisation usage is aggregated into StageResult.NormalizationUsage
// and never counted as solution cost.
func (a *Agent) RunBaseline(ctx context.Context, contracts []Contract) (scoring.StageResult, error) {
	res := scoring.StageResult{Stage: "s0_baseline", Model: a.Client.Model()}
	var normUsage scoring.Usage
	for _, c := range contracts {
		start := time.Now()
		free, u1, err := a.Baseline(ctx, c.Text)
		if err != nil {
			return res, fmt.Errorf("baseline %s: %w", c.ID, err)
		}
		findings, u2, err := a.Normalize(ctx, free)
		if err != nil {
			// A weak baseline that cannot emit structured output scores as
			// "no findings" rather than aborting the whole run.
			fmt.Fprintf(os.Stderr, "warn: normalize %s: %v (treating as no findings)\n", c.ID, err)
			findings = nil
		}
		normUsage.Add(toScoringUsage(u2))
		res.Contracts = append(res.Contracts, scoring.ContractResult{
			ContractID: c.ID,
			Findings:   findings,
			Usage:      toScoringUsage(u1), // solution cost = the review call only
			DurationMs: durMs(start),
		})
	}
	res.NormalizationUsage = &normUsage
	return res, nil
}

// RunStructured runs S1 (whole document) or S2 (per section) over the corpus.
func (a *Agent) RunStructured(ctx context.Context, contracts []Contract, perSection bool) (scoring.StageResult, error) {
	stage := "s1_structured"
	if perSection {
		stage = "s2_sections"
	}
	res := scoring.StageResult{Stage: stage, Model: a.Client.Model()}
	for _, c := range contracts {
		start := time.Now()
		var (
			findings []scoring.Finding
			u        llm.Usage
			err      error
		)
		if perSection {
			findings, u, err = a.StructuredPerSection(ctx, c.Text)
		} else {
			findings, u, err = a.Structured(ctx, c.Text)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "warn: %s %s: %v (treating as no findings)\n", stage, c.ID, err)
			findings = nil
		}
		res.Contracts = append(res.Contracts, scoring.ContractResult{
			ContractID: c.ID, Findings: findings, Usage: toScoringUsage(u), DurationMs: durMs(start),
		})
	}
	return res, nil
}

// stageName maps a short stage key to its result-file stage label.
var stageName = map[string]string{
	"s3": "s3_checklist",
	"s4": "s4_absence",
	"s5": "s5_gated",
}

// RunPipeline runs S3, S4, or S5 over the corpus. S3 is checklist only; S4 adds
// the absence pass; S5 additionally applies the citation gate to clause-based
// findings. Per-contract trajectories are written to trajDir when non-empty.
func (a *Agent) RunPipeline(ctx context.Context, contracts []Contract, stage, trajDir string) (scoring.StageResult, error) {
	label, ok := stageName[stage]
	if !ok {
		return scoring.StageResult{}, fmt.Errorf("unknown stage %q (want s3, s4, or s5)", stage)
	}
	res := scoring.StageResult{Stage: label, Model: a.Client.Model()}
	for _, c := range contracts {
		start := time.Now()
		var usage llm.Usage

		checklist, u1, traj1, err := a.Checklist(ctx, c.Text)
		if err != nil {
			return res, fmt.Errorf("%s checklist %s: %w", label, c.ID, err)
		}
		addUsage(&usage, u1)
		steps := traj1

		var absent []scoring.Finding
		if stage == "s4" || stage == "s5" {
			ab, u2, traj2, err := a.Absence(ctx, c.Text)
			if err != nil {
				return res, fmt.Errorf("%s absence %s: %w", label, c.ID, err)
			}
			addUsage(&usage, u2)
			absent = ab
			steps = append(steps, traj2...)
		}

		var final []scoring.Finding
		var rejected []Rejected
		if stage == "s5" {
			kept, rej := a.CitationGate(c.Text, checklist)
			final = append(kept, absent...)
			rejected = rej
		} else {
			final = append(checklist, absent...)
		}

		if trajDir != "" {
			md := RenderTrajectory(label, c.ID, steps, final, rejected)
			_ = os.WriteFile(filepath.Join(trajDir, fmt.Sprintf("%s_%s.md", label, c.ID)), []byte(md), 0o644)
		}

		res.Contracts = append(res.Contracts, scoring.ContractResult{
			ContractID: c.ID, Findings: final, Usage: toScoringUsage(usage), DurationMs: durMs(start),
		})
	}
	return res, nil
}

// WriteStageResult marshals a stage result to path as indented JSON.
func WriteStageResult(path string, res scoring.StageResult) error {
	b, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, 0o644)
}

// ReadStageResult loads a stage result JSON file.
func ReadStageResult(path string) (scoring.StageResult, error) {
	var res scoring.StageResult
	b, err := os.ReadFile(path)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(b, &res)
	return res, err
}

// LoadCorpus is a thin convenience wrapper used by the cmd runners.
func LoadCorpus(path string) (*statutes.Corpus, error) { return statutes.Load(path) }
