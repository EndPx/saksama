// Command eval scores two stage-result files against the ground truth and
// writes a markdown comparison table. It calls no LLM, so judges can verify
// every committed number without an API key.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/EndPx/saksama/internal/agent"
	"github.com/EndPx/saksama/internal/scoring"
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
