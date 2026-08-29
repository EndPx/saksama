// Command solution runs the staged solution reviewers (S1-S5) over the corpus,
// and can also render the final memos from an existing result file.
//
//	solution -stage all        # run S2..S5 in sequence (needs SAKSAMA_* env)
//	solution -stage s3         # run one stage
//	solution -memos -from results/s5_gated.json   # render memos, no API key
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/EndPx/saksama/internal/agent"
	"github.com/EndPx/saksama/internal/llm"
	"github.com/EndPx/saksama/internal/memo"
	"github.com/EndPx/saksama/internal/scoring"
	"github.com/EndPx/saksama/internal/statutes"
)

func main() {
	stage := flag.String("stage", "all", "stage to run: s1, s2, s3, s4, s5, or all (S2..S5)")
	statutesPath := flag.String("statutes", "data/statutes/2026-08.yaml", "statute corpus")
	contractsDir := flag.String("contracts", "data/contracts", "contracts directory")
	resultsDir := flag.String("results", "results", "results output directory")
	trajDir := flag.String("traj", "trajectories", "trajectory output directory")
	memos := flag.Bool("memos", false, "render memos from -from result file (no API key needed)")
	from := flag.String("from", "results/s5_gated.json", "stage-result file to render memos from")
	memosDir := flag.String("memos-out", "memos", "memo output directory")
	reviewFile := flag.String("contract", "", "review ONE plain-text contract file (live; needs SAKSAMA_* env)")
	demo := flag.Bool("demo", false, "keyless demo: summarise committed S5 results for a few contracts")
	flag.Parse()

	_ = llm.LoadEnvFile(".env") // auto-load .env if present (shell env still wins)

	corpus, err := agent.LoadCorpus(*statutesPath)
	must(err)

	if *demo {
		runDemo(corpus, *from, *contractsDir)
		return
	}
	if *reviewFile != "" {
		reviewOne(corpus, *reviewFile)
		return
	}

	contracts, err := agent.LoadContracts(*contractsDir)
	must(err)

	if *memos {
		renderMemos(*from, contracts, *memosDir, corpus)
		return
	}

	client, err := llm.New()
	must(err)
	a := agent.New(client, corpus)
	ctx := context.Background()

	run := func(s string) {
		res, err := runStage(ctx, a, contracts, s, *trajDir)
		must(err)
		out := filepath.Join(*resultsDir, resultFile(s))
		must(agent.WriteStageResult(out, res))
		fmt.Fprintf(os.Stderr, "wrote %s\n", out)
	}

	if *stage == "all" {
		for _, s := range []string{"s2", "s3", "s4", "s5"} {
			run(s)
		}
		return
	}
	run(*stage)
}

func runStage(ctx context.Context, a *agent.Agent, contracts []agent.Contract, stage, trajDir string) (scoring.StageResult, error) {
	switch stage {
	case "s1":
		return a.RunStructured(ctx, contracts, false)
	case "s2":
		return a.RunStructured(ctx, contracts, true)
	case "s3", "s4", "s5":
		return a.RunPipeline(ctx, contracts, stage, trajDir)
	default:
		return scoring.StageResult{}, fmt.Errorf("unknown stage %q", stage)
	}
}

func resultFile(stage string) string {
	switch stage {
	case "s1":
		return "s1_structured.json"
	case "s2":
		return "s2_sections.json"
	case "s3":
		return "s3_checklist.json"
	case "s4":
		return "s4_absence.json"
	case "s5":
		return "s5_gated.json"
	}
	return stage + ".json"
}

func renderMemos(from string, contracts []agent.Contract, outDir string, corpus *statutes.Corpus) {
	res, err := agent.ReadStageResult(from)
	must(err)
	text := map[string]string{}
	for _, c := range contracts {
		text[c.ID] = c.Text
	}
	for _, cr := range res.Contracts {
		md := memo.Render(cr.ContractID, text[cr.ContractID], "", cr.Findings, corpus)
		out := filepath.Join(outDir, cr.ContractID+".md")
		must(os.WriteFile(out, []byte(md), 0o644))
		fmt.Fprintf(os.Stderr, "wrote %s\n", out)
	}
}

// reviewOne runs the live S5 engine on a single plain-text contract and prints
// the triage memo to stdout. This is generation, so it needs the SAKSAMA_* env.
func reviewOne(corpus *statutes.Corpus, path string) {
	b, err := os.ReadFile(path)
	must(err)
	text := string(b)
	client, err := llm.New()
	must(err)
	a := agent.New(client, corpus)
	res, err := a.RunPipeline(context.Background(), []agent.Contract{{ID: filepath.Base(path), Text: text}}, "s5", "")
	must(err)
	fmt.Print(memo.Render(filepath.Base(path), text, "", res.Contracts[0].Findings, corpus))
}

// runDemo prints a compact, KEYLESS summary of a few contracts from the committed
// S5 results — no LLM call, so a judge sees the product output with no API key.
func runDemo(corpus *statutes.Corpus, resultsFile, contractsDir string) {
	res, err := agent.ReadStageResult(resultsFile)
	must(err)
	byID := map[string]scoring.ContractResult{}
	for _, c := range res.Contracts {
		byID[c.ContractID] = c
	}
	demos := []struct{ ID, Label, Take string }{
		{"c02", "Clean PKWT", "No false positives — nothing flagged, no missing protection invented"},
		{"c01", "Problematic PKWT", "Findings carry verbatim evidence + legal basis"},
		{"c04", "Problematic PKWTT", "Contract-type aware — PKWT-only rules skipped; caught the PKWTT probation limit"},
	}
	fmt.Println("==================================================")
	fmt.Println(" SAKSAMA DEMO  (keyless, from committed S5 results)")
	fmt.Println("==================================================")
	for i, d := range demos {
		cr := byID[d.ID]
		b, _ := os.ReadFile(filepath.Join(contractsDir, d.ID, "contract.md"))
		typ := "PKWT"
		if strings.Contains(strings.ToUpper(string(b)), "PKWTT") {
			typ = "PKWTT"
		}
		var legal, missing int
		for _, f := range cr.Findings {
			if f.Deteksi == statutes.DeteksiTidakAdaKlausa {
				missing++
			} else {
				legal++
			}
		}
		fmt.Printf("\n[%d/%d] %s  (type: %s)\n", i+1, len(demos), d.Label, typ)
		fmt.Printf("    Legal findings: %d    Missing protections: %d    Total: %d\n", legal, missing, len(cr.Findings))
		for _, f := range cr.Findings {
			if f.Tier == statutes.TierBatalDemiHukum {
				if p, ok := corpus.Get(f.StatuteID); ok {
					fmt.Printf("    Critical: %s (%s %s)\n", p.Judul, p.DasarHukum, p.Pasal)
				}
				break
			}
		}
		fmt.Printf("    OK: %s\n", d.Take)
	}
	fmt.Println("\nThis is the keyless benchmark output. For a LIVE review of any contract")
	fmt.Println("(needs an API key in .env):")
	fmt.Println("  go run ./cmd/solution -contract examples/pkwt-problematic.txt")
}

func must(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
