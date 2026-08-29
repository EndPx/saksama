# Saksama — build and run targets.
# baseline/solution need the SAKSAMA_* environment (they call the LLM).
# eval and memos run with NO API key, from committed result files.

GO ?= go

.PHONY: build test baseline solution eval memos trajectories all clean

build:
	$(GO) build ./...

test:
	$(GO) test ./...

# S0 baseline: one call per contract + a separate normalisation call.
baseline:
	$(GO) run ./cmd/baseline

# S1..S5 in sequence (writes results/s*.json and trajectories/).
solution:
	$(GO) run ./cmd/solution -stage all

# Score baseline vs final solution and write results/comparison.md. No API key.
eval:
	$(GO) run ./cmd/eval

# Render the triage memos from the committed S5 result. No API key.
memos:
	$(GO) run ./cmd/solution -memos -from results/s5_gated.json

# Export Claude Code session trajectories to trajectories/.
trajectories:
	$(GO) run ./cmd/trajectory

# Full pipeline (needs API key for baseline+solution).
all: baseline solution eval memos

clean:
	rm -f results/s0_baseline.json results/s1_structured.json results/s2_sections.json \
	      results/s3_checklist.json results/s4_absence.json results/s5_gated.json \
	      results/comparison.md memos/*.md
