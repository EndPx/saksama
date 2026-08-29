# Saksama

**An evidence-first agentic review engine for Indonesian employment contracts that
identifies applicable risks and missing mandatory protections — while minimizing
false positives.**

*Saksama* is Bahasa Indonesia for *careful, thorough, meticulous scrutiny.*

---

## The problem

A fresh graduate in Indonesia who has just been handed a PKWT or PKWTT contract has a
day or two to decide whether to sign, usually with no lawyer to call. The trap is
rarely the scary-looking clause that *is* in the document — it is the mandatory
protection that is *absent* from it (a probation period a PKWT may not contain at all,
compensation the employer owes but never mentions, overtime rights quietly waived). An
unstructured language model is reactive: it comments on the text in front of it, so it
systematically misses what is missing.

## What Saksama does

It reads the contract, classifies its type, checks it against **14 verified provisions**
of Indonesian labour law, and produces a **triage memo** (in Indonesian) with findings
graded by legal severity, the mandatory protections that are missing, concrete questions
to ask HR, and a marker for when to seek real legal help. It is a triage aid, **not a
verdict and not legal advice** — every memo says so.

## Why Saksama?

Most contract review focuses on finding suspicious text. Saksama additionally asks, for
every finding:

1. **Does this rule actually apply to this contract type?** (PKWT rules are not applied
   to a PKWTT contract; the overtime-exemption rule only fires on a genuine waiver.)
2. **Is a required protection absent?** (an explicit pass asks what is missing.)
3. **Is there evidence in the contract supporting the finding?** (every clause finding
   must quote the contract verbatim, verified in Go.)
4. **Is the finding strong enough to report?** (a conservative checklist, so a false
   positive doesn't needlessly alarm someone with no legal background.)

Applicability and evidence are first-class constraints. This is why the precision work
matters: on the committed evaluation, **zero false positives occur on clean contracts.**

## Quick Start

**Requirements:** Go 1.22+ only. **No API key** is required to reproduce the committed
evaluation. **Make is optional** — every `make` target below is a one-line wrapper, so
on Windows (or anywhere without Make) just use the `go run` form.

```bash
git clone https://github.com/EndPx/saksama.git
cd saksama

go run ./cmd/eval      # or, if you have Make:  make eval
```

This runs the **frozen, keyless, network-free** evaluation over the committed
`results/*.json`. Expected output:

```
Recall (primary)    100.0%
Precision            72.2%
F1                   83.9%
False-positive rate  27.8%
Absence detection    100.0%
False positives on clean   0
```

Other commands:

```bash
go test ./...                 # unit + audit regression tests
go build ./...                # build all binaries
cat results/comparison.md     # baseline vs S5 table
cat results/audit.md          # confusion matrix + citation grounding
cat docs/FAILURE_MODES.md     # every remaining false positive, classified
```

### Review a single contract

Keyless — a compact review summary for a few contracts, straight from the committed
results (no API key, no network):

```bash
go run ./cmd/solution -demo      # or, with Make:  make demo
```

Live — review any plain-text contract with the engine (this is *generation*, so it
needs an API key; it is not the keyless evaluation):

```bash
go run ./cmd/solution -contract examples/pkwt-problematic.txt
```

The engine is **format-agnostic — it reviews contract text.** PDF/DOCX ingestion is a
separate input adapter (roadmap), intentionally outside the frozen evaluation engine.
Sample contracts live in [`examples/`](examples/).

### Run with Docker

The image carries **no API key** and its default entrypoint is the keyless evaluation:

```bash
docker build -t saksama .

# Keyless: reproduce the committed evaluation (default entrypoint)
docker run --rm saksama

# Keyless: the compact demo summary
docker run --rm --entrypoint /app/bin/solution saksama -demo

# Live review of a contract (needs a key; --env-file supplies SAKSAMA_*)
docker run --rm --env-file .env --entrypoint /app/bin/solution \
  saksama -contract examples/pkwt-problematic.txt
```

The build is multi-stage (Go 1.22 → distroless static), so the runtime image ships only
the binaries and the committed data/results — no toolchain, no key.

## Example (excerpt from a generated memo)

From [`memos/c01.md`](memos/c01.md):

```
### Batal demi hukum

**PKWT dilarang mensyaratkan masa percobaan** (Pasal 4)

> PIHAK KEDUA wajib menjalani masa percobaan kerja selama 3 (tiga) bulan
> terhitung sejak tanggal mulai bekerja. ...

Dasar hukum: PP 35/2021 Pasal 12. Klausul masa percobaan 3 bulan dalam PKWT
melanggar PP 35/2021 Pasal 12 karena PKWT dilarang mensyaratkan masa
percobaan. Klausul ini batal demi hukum dan masa kerja dihitung sejak awal.
Tingkat keyakinan penilaian: A.
```

(`(Pasal 4)` refers to the article of the *contract* under review; `Pasal 12` is
the statute. Quoted verbatim from [`memos/c01.md`](memos/c01.md).)

Each finding carries a verbatim contract quote, the legal basis, what it means for the
worker, and a confidence level.

## Architecture

```
Contract
   |
   v
Contract-type classification (PKWT / PKWTT)
   |
   v
Structured extraction (parse into articles)
   |
   +----------------+
   v                v
Clause checks    Absence checks (what is missing)
   |                |
   +-------+--------+
           v
     Rule applicability (skip rules that don't apply to this type)
           v
     Evidence / citation gate (verbatim quote, verified in Go)
           v
     Severity + deterministic scoring
           v
        Triage memo
```

The build is split into **LLM-dependent** stages (`internal/agent`, `internal/llm`) and
a **deterministic** scorer (`internal/scoring`, `cmd/eval`). Full diagram and the
determinism boundary: [`docs/ARCHITECTURE.md`](docs/ARCHITECTURE.md).

## Agentic workflow

Built and measured in stages, so the improvement is visible at every step:

```
Baseline  ->  S1 Structured  ->  S2 Sections  ->  S3 Checklist  ->  S4 Absence  ->  S5 Evidence Gate
```

| Stage | Recall | Precision | F1 | Absence |
|-------|-------:|----------:|---:|--------:|
| S0 baseline | 23.1% | 3.7% | 6.3% | 33.3% |
| S1 structured | 46.2% | 50.0% | 48.0% | 16.7% |
| **S2 sections** | **38.5%** | **18.5%** | **25.0%** | 50.0% |
| S3 checklist | 53.8% | 63.6% | 58.3% | 0.0% |
| S4 absence | 100% | 65.0% | 78.8% | 100% |
| S5 evidence gate | 100% | 72.2% | 83.9% | 100% |

**S2 is kept as a failed experiment**: splitting the contract and reviewing one article
at a time *dropped* both recall and precision below S1, because a section read in
isolation loses cross-article context. It demonstrates that simply adding another
agentic stage does not automatically improve quality — the opposite of the usual
"more agents = better" assumption. Full per-stage reasoning is in
[`CHANGELOG.md`](CHANGELOG.md).

## Evaluation & results

Measured over **12 synthetic contracts (13 ground-truth findings)**, model
`minimax/minimax-m3:free`, full run ~29 min at **\$0**.

> These are **benchmark results on the committed synthetic evaluation corpus**. They
> are not a claim of general accuracy on real Indonesian employment contracts — 100%
> recall here means "found all 13 planted findings on these 12 contracts," not
> "Saksama is 100% accurate." The corpus size (12) is stated openly for exactly this
> reason.

**Finding-level metrics** (each ground-truth finding counted individually; the nine
PP35-13 sub-checks count separately). False-positive rate = false positives ÷ reported
findings = 1 − precision.

| Metric | Baseline (S0) | Final (S5) |
|--------|--------------:|-----------:|
| Recall | 23.1% | **100%** |
| Precision | 3.7% | **72.2%** |
| F1 | 6.3% | **83.9%** |
| False-positive rate | 96.3% | **27.8%** |
| Absence detection | 33.3% | **100%** |
| Tier accuracy | 33.3% | **100%** |
| Cross-section recall | 100% | **100%** |
| Citation presence accuracy | 18.8% | **100%** |
| Citation location accuracy | 9.4% | **100%** |
| Reported findings | 82 | 18 |
| False positives | 79 | **5** |
| False positives on clean contracts | 0 | **0** |

**Cell-level confusion matrix (S5)** — a *different unit of evaluation* from the
finding-level table above: one decision per (contract × statute) pair, 168 cells, with
the nine PP35-13 sub-checks collapsed into one cell. This is why cell-level precision
(66.7%) differs slightly from finding-level precision (72.2%): same 5 false positives,
different denominator.

| | Ground-truth Positive | Ground-truth Negative |
|---|---:|---:|
| **Predicted Positive** | 10 (TP) | 5 (FP) |
| **Predicted Negative** | 0 (FN) | 153 (TN) |

Source artifacts: [`results/comparison.md`](results/comparison.md),
[`results/audit.md`](results/audit.md).

**Citation accuracy is deterministic and syntactic + location-based, NOT semantic:** it
verifies the quoted excerpt exists in the contract and lies within the cited article. It
does **not** verify that the excerpt semantically supports the finding — that would need
an LLM judge and is documented as future work, not claimed.

## Known failure modes / limitations

The final evaluation has **5 false positives**:

- 3 interpretation / adjacent-rule over-application
- 1 extraction error
- 1 genuinely ambiguous case

Importantly: **no false positives occur on clean contracts**, every remaining false
positive lands on genuinely risky text (the engine over-tags an adjacent rule, it does
not hallucinate on compliant clauses), and **no known scoring defect remains** (the Go
metrics were independently re-derived in Python and matched exactly). These are recorded
as known limitations rather than silently patched. Other limits: the corpus is 12
synthetic contracts; the free model is non-deterministic, so a fresh *generation* may
differ by a few findings (the committed `results/` are the frozen baseline). Full
per-FP audit: [`docs/FAILURE_MODES.md`](docs/FAILURE_MODES.md) +
[`results/failure_modes.json`](results/failure_modes.json).

## Reproducibility

The committed `results/` and `memos/` are the frozen evaluation output used for this
submission. To reproduce the scoring from scratch:

```bash
make eval        # or: go run ./cmd/eval
```

The evaluation is **keyless, network-free, deterministic, and independently auditable.**
Two consecutive `cmd/eval` runs produce identical hashes for `results/comparison.md` and
`results/audit.md`. The scoring was independently re-derived in Python and matched the Go
implementation. Full instructions, runtime, and cost: [`REPRODUCTION.md`](REPRODUCTION.md).

## Optional: model-backed generation

To *generate a new review* instead of reproducing the frozen evaluation, configure an
OpenAI-compatible provider (e.g. OpenRouter) and run the LLM stages:

```bash
cp .env.example .env      # set SAKSAMA_MODEL / SAKSAMA_API_BASE / SAKSAMA_API_KEY
set -a; . ./.env; set +a  # load env (bash/zsh; on Windows set the vars your shell's way)
go run ./cmd/baseline
go run ./cmd/solution -stage all
go run ./cmd/solution -memos -from results/s5_gated.json
```

Note: generation is **non-deterministic** (the free `minimax-m3` path varies even at
temperature 0), so freshly generated findings may differ slightly from the committed
artifacts. The **evaluation is deterministic; generation is not** — the two are
deliberately separated (see [`docs/ARCHITECTURE.md`](docs/ARCHITECTURE.md)).

## Security & Safety

- No API keys, tokens, credentials, or real contract data are committed.
- Evaluation is keyless and network-free.
- Runtime LLM access uses environment-provided credentials only.
- All evaluation contracts are synthetic.
- Saksama produces a triage aid, not legal advice; the final decision remains with the user.

See [`SECURITY.md`](SECURITY.md) for the full security model and credential-handling notes.

## Project structure

```
cmd/            baseline, solution, eval, trajectory  (entry points)
internal/       llm, agent, contract, statutes, scoring, memo
data/           statutes/ (14 provisions) + contracts/ (12 synthetic + ground truth)
examples/       synthetic plain-text contracts for the review CLI / make demo
results/        frozen evaluation JSON, comparison.md, audit.md, failure_modes.json
memos/          12 generated triage memos
trajectories/   per-contract agent trajectories (S3-S5)
docs/           ARCHITECTURE.md, FAILURE_MODES.md
```

Specification: [`CLAUDE.md`](CLAUDE.md). Everything here was built during the micro1
Agentic Workflows Hackathon; nothing existed before.

## Roadmap

The current submission is the backend review engine (frozen). Planned product-layer work,
intentionally outside the evaluation engine:

- PDF / DOCX ingestion — a text-extraction *input adapter* feeding the format-agnostic
  engine (the engine already accepts contract text; ingestion is orthogonal to it)
- Contract Coverage Map
- Negotiation Mode
- HR-question copy actions
- Unknown contract-type UX
