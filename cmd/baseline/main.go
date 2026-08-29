// Command baseline runs the single-call baseline reviewer over the corpus and
// writes results/s0_baseline.json. Requires the SAKSAMA_* environment.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/EndPx/saksama/internal/agent"
	"github.com/EndPx/saksama/internal/llm"
)

func main() {
	statutesPath := flag.String("statutes", "data/statutes/2026-08.yaml", "statute corpus")
	contractsDir := flag.String("contracts", "data/contracts", "contracts directory")
	out := flag.String("out", "results/s0_baseline.json", "output stage-result file")
	flag.Parse()

	corpus, err := agent.LoadCorpus(*statutesPath)
	must(err)
	contracts, err := agent.LoadContracts(*contractsDir)
	must(err)
	_ = llm.LoadEnvFile(".env") // auto-load .env if present (shell env still wins)
	client, err := llm.New()
	must(err)

	a := agent.New(client, corpus)
	res, err := a.RunBaseline(context.Background(), contracts)
	must(err)
	must(agent.WriteStageResult(*out, res))
	fmt.Fprintf(os.Stderr, "wrote %s (%d contracts)\n", *out, len(res.Contracts))
}

func must(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
