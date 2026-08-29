# Reproduction

Written for someone on a clean machine.

## Toolchain

- Go 1.22 or newer (tested on 1.27)
- Optional: Docker (for the hermetic build path)

## Model / provider

The runs in this repo used **`minimax/minimax-m3:free` via OpenRouter**, an
OpenAI-compatible Chat Completions endpoint. `internal/llm` speaks that
`/chat/completions` shape (Authorization: Bearer, messages-with-system,
`choices[].message.content`). This is a deliberate deviation from CLAUDE.md's
"Anthropic Messages API" wording — the client is still hand-written `net/http`
with no SDK. Any OpenAI-compatible endpoint works; set the three variables below.

## Environment variables

Copy `.env.example` to `.env` and fill in:

- `SAKSAMA_MODEL` — e.g. `minimax/minimax-m3:free`
- `SAKSAMA_API_BASE` — e.g. `https://openrouter.ai/api/v1`
- `SAKSAMA_API_KEY` — your key (never committed)
- `SAKSAMA_PRICE_IN` / `SAKSAMA_PRICE_OUT` — optional USD per million tokens; leave
  empty to report cost as 0 (the free model above costs \$0)

Load them before running: `set -a; . ./.env; set +a`.

## Run everything

```bash
go build -o bin/baseline ./cmd/baseline
go build -o bin/solution ./cmd/solution
go build -o bin/eval     ./cmd/eval

./bin/baseline                 # S0  -> results/s0_baseline.json
./bin/solution -stage all      # S2..S5 -> results/s*.json + trajectories/
./bin/solution -stage s1       # optional S1
./bin/eval                     # -> results/comparison.md
./bin/solution -memos -from results/s5_gated.json   # -> memos/*.md
```

Or use the Makefile: `make baseline`, `make solution`, `make eval`, `make memos`.

## Verify results without an API key

`results/` and `memos/` are committed. Judges can confirm every number with no key
and no network:

```bash
go run ./cmd/eval        # reads results/*.json, recomputes the table, calls no LLM
```

## Measured runtime and cost

Full run (baseline + S1–S5 + eval + memos, 12 contracts, ~630 model calls) on
`minimax/minimax-m3:free`: **~29 minutes wall-clock, total cost \$0** (free model).
Average time per contract was ~32 s at baseline and ~38 s at S5. A paid fast model
would finish in a few minutes; a reasoning model (e.g. Nemotron) took ~150 s/call and
is not recommended — see `CHANGELOG.md`, "Model selection".

## Expected output

`go run ./cmd/eval` prints and writes `results/comparison.md`:

```
| Metric | s0_baseline | s5_gated |
|--------|------|------|
| Recall (primary) | 23.1% | 100.0% |
| Precision | 3.7% | 72.2% |
| F1 | 6.3% | 83.9% |
| False-positive rate | 96.3% | 27.8% |
| Absence detection rate | 33.3% | 100.0% |
| Tier accuracy | 33.3% | 100.0% |
| Cross-section recall | 100.0% | 100.0% |
| Recall (confidence A only) | 27.3% | 100.0% |
| Correct / reported | 3 / 82 | 13 / 18 |
| False positives on clean | 0 | 0 |
```

Note: the model is not perfectly deterministic even at temperature 0, so a fresh run
may differ by a few findings. The committed `results/` are the run described in
`CHANGELOG.md`.
