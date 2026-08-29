# Reproduction

Written for someone on a clean machine. Numbers and exact timings are filled in as stages
land; this file tracks the intended, tested path.

## Toolchain

- Go 1.22 or newer
- Docker (for the hermetic build path)

## Environment variables

Copy `.env.example` to `.env` and fill in:

- `SAKSAMA_MODEL` — the model id (never hardcoded in the code)
- `SAKSAMA_API_BASE` — the Messages API endpoint
- `SAKSAMA_API_KEY` — your key (never committed)

## Run everything

```bash
make baseline   # S0 — one LLM call per contract
make solution   # S2..S5 in sequence
make eval       # deterministic scoring + results/comparison.md
make memos      # render the triage memos
```

## Verify without an API key

The `results/` and `memos/` directories are committed. Judges can confirm every number
without a key or a network call:

```bash
make eval       # reads existing results/*.json, recomputes the comparison table, calls no LLM
```

## Expected runtime and cost

_pending — filled in after S0 lands with real token-usage numbers from the API._
