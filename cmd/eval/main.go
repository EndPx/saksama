// Command eval scores two stage-result files against the ground truth and
// writes a markdown comparison table. It calls no LLM, so judges can verify
// every committed number without an API key.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/EndPx/saksama/internal/agent"
	"github.com/EndPx/saksama/internal/scoring"
	"github.com/EndPx/saksama/internal/statutes"
)

func main() {
	base := flag.String("base", "results/s0_baseline.json", "baseline stage-result file")
	sol := flag.String("sol", "results/s5_gated.json", "solution stage-result file")
	statutesPath := flag.String("statutes", "data/statutes/2026-08.yaml", "statute corpus")
	contractsDir := flag.String("contracts", "data/contracts", "contracts directory (for ground truth)")
	out := flag.String("out", "results/comparison.md", "comparison output file")
	flag.Parse()

	corpus, err := agent.LoadCorpus(*statutesPath)
	must(err)

	truths, err := loadTruths(*contractsDir)
	must(err)

	baseRes, err := agent.ReadStageResult(*base)
	must(err)
	solRes, err := agent.ReadStageResult(*sol)
	must(err)

	baseM := scoring.Evaluate(corpus, truths, baseRes)
	solM := scoring.Evaluate(corpus, truths, solRes)

	table := scoring.Compare(baseM, solM)
	fmt.Print(table)
	must(os.WriteFile(*out, []byte(table), 0o644))
	fmt.Fprintf(os.Stderr, "\nwrote %s\n", *out)

	// Audit artifacts: cell-level confusion matrix + citation grounding.
	contracts, err := agent.LoadContracts(*contractsDir)
	must(err)
	textByID := map[string]string{}
	for _, c := range contracts {
		textByID[c.ID] = c.Text
	}
	audit := renderAudit(corpus, truths, textByID, baseRes, solRes)
	must(os.WriteFile("results/audit.md", []byte(audit), 0o644))
	fmt.Fprintln(os.Stderr, "wrote results/audit.md")
}

func renderAudit(corpus *statutes.Corpus, truths map[string]scoring.Truth, text map[string]string, base, sol scoring.StageResult) string {
	cmBase := scoring.ConfusionMatrix(corpus, truths, base)
	cmSol := scoring.ConfusionMatrix(corpus, truths, sol)
	ctBase := scoring.Citation(text, base)
	ctSol := scoring.Citation(text, sol)
	pct := func(f float64) string { return fmt.Sprintf("%.1f%%", f*100) }
	matrix := func(name string, c scoring.Confusion) string {
		return fmt.Sprintf(
			"### %s\n\n```\n                 Ground Truth\n              Positive   Negative\nPred Positive   %4d       %4d\nPred Negative   %4d       %4d\n```\n\nCell-level precision %s, recall %s (over %d contract x statute cells).\n\n",
			name, c.TP, c.FP, c.FN, c.TN, pct(c.Precision()), pct(c.Recall()), c.TP+c.FP+c.FN+c.TN)
	}
	var b strings.Builder
	b.WriteString("# Evaluation audit\n\n")
	b.WriteString("Cell-level confusion matrix over (contract x statute) decisions. This exposes\n")
	b.WriteString("true negatives; it is a coarser granularity than the finding-level headline\n")
	b.WriteString("metrics (the nine PP35-13 sub-checks collapse to one cell here).\n\n")
	b.WriteString("## Confusion matrix\n\n")
	b.WriteString(matrix("Baseline (S0)", cmBase))
	b.WriteString(matrix("Solution (S5)", cmSol))
	b.WriteString("## Citation grounding\n\n")
	b.WriteString("Deterministic and syntactic/location-based, NOT semantic (see internal/scoring/audit.go).\n\n")
	b.WriteString("| Metric | Baseline (S0) | Solution (S5) |\n|---|---|---|\n")
	fmt.Fprintf(&b, "| Clause-based findings | %d | %d |\n", ctBase.ClauseFindings, ctSol.ClauseFindings)
	fmt.Fprintf(&b, "| With a non-empty quote | %d | %d |\n", ctBase.WithQuote, ctSol.WithQuote)
	fmt.Fprintf(&b, "| Quote present in contract (presence accuracy) | %s | %s |\n", pct(ctBase.PresenceAccuracy()), pct(ctSol.PresenceAccuracy()))
	fmt.Fprintf(&b, "| Quote within cited article (location accuracy) | %s | %s |\n", pct(ctBase.LocationAccuracy()), pct(ctSol.LocationAccuracy()))
	b.WriteString("\nSemantic support of the quote for the finding is not verified deterministically; see `docs/FAILURE_MODES.md`.\n")
	return b.String()
}

func loadTruths(dir string) (map[string]scoring.Truth, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	truths := map[string]scoring.Truth{}
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		p := filepath.Join(dir, e.Name(), "truth.yaml")
		if _, err := os.Stat(p); err != nil {
			continue
		}
		t, err := scoring.LoadTruth(p)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", p, err)
		}
		truths[e.Name()] = t
	}
	return truths, nil
}

func must(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
