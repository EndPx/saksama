# Saksama

**Saksama** (Bahasa Indonesia: *careful, thorough, meticulous scrutiny*) is an agent that
reviews Indonesian employment contracts — PKWT and PKWTT — for the person with the least
leverage in the room: a fresh graduate who has just been handed an offer and a two- or
three-day window to sign, and no lawyer to call.

## The user and the bottleneck

A fresh graduate reading an employment contract cannot tell which clauses are merely
unfavourable and which are actually void under Indonesian labour law. Worse, the thing that
most often traps entry-level workers is not a frightening clause that *is* in the document —
it is a mandatory protection that is *absent* from it. A probation period that should not
exist in a PKWT at all; compensation money the employer is required to pay but never
mentions; overtime rights quietly waived without defining the senior job category that would
make the waiver lawful. An unstructured language model is reactive: it comments on the text
in front of it, so it systematically misses what is missing.

Saksama attacks absence directly. It walks a checklist grounded in fourteen verified
provisions of Indonesian law and forces the question *"what should be here that isn't?"* for
every one of them, then grades what it finds by legal severity — from void-by-law down to
non-binding policy guidance — and cites the contract article and the statute for each claim.

## What it produces

Not a verdict. A **triage memo**, in Indonesian, addressed to the worker: findings grouped by
severity, the mandatory protections that are missing, a set of polite questions to send HR
before signing, and a marker for when the stakes warrant real legal help. The memo states
plainly that it is an automated triage aid, not legal advice, and that the final decision
rests with the reader.

## Results

Measured over 12 synthetic contracts (13 ground-truth findings), model
`minimax/minimax-m3:free` via OpenRouter, full run ~29 min at \$0:

| Metric | Baseline (S0) | Solution (S5) |
|---|---|---|
| Recall | 23.1% | **100%** |
| Precision | 3.7% | **72.2%** |
| F1 | 6.3% | **83.9%** |
| Absence detection | 33.3% | **100%** |
| Tier accuracy | 33.3% | **100%** |
| Cross-section recall | 100% | **100%** |
| False positives on clean | 0 | **0** |

Absence detection stays low (17–50%) through structured output, per-section reading, and a
clause-by-clause checklist (0%) — and only reaches 100% once an explicit pass forces the agent
to ask what is *missing* (S4/S5). The full staged progression — including one experiment
(per-section S2) that made results worse, and a precision-refinement iteration that lifted S5
precision from 27% to 72% by classifying the contract type before applying rules — is in
[CHANGELOG.md](CHANGELOG.md); the specification is in [CLAUDE.md](CLAUDE.md); how to run and
verify keyless is in [REPRODUCTION.md](REPRODUCTION.md).

## Audit & artifacts

Everything is committed so the evaluation can be inspected without running anything:

- [`results/comparison.md`](results/comparison.md) — baseline vs S5 metric table.
- [`results/audit.md`](results/audit.md) — cell-level confusion matrix and citation
  grounding (presence + location accuracy; deterministic, syntactic — not semantic).
- [`docs/FAILURE_MODES.md`](docs/FAILURE_MODES.md) + [`results/failure_modes.json`](results/failure_modes.json)
  — every remaining false positive, classified, kept visible rather than hidden.
- [`docs/ARCHITECTURE.md`](docs/ARCHITECTURE.md) — pipeline diagram and the boundary
  between LLM-dependent stages and the deterministic scorer.
- [`SECURITY.md`](SECURITY.md) — zero secrets, synthetic data, human-in-the-loop.

`cmd/eval` is deterministic (two runs produce identical output hashes) and makes no
network calls, so judges verify every number with `make eval` and no API key.

## Architecture (planned)

A hand-written Go agent loop — no orchestration framework, `net/http` straight to the
Messages API — run in measured stages so the improvement over the baseline is visible at
every step:

- **Baseline** — one LLM call, whole contract, free-form "list the risks".
- **S1 structured** → **S2 section parser** → **S3 checklist-driven** → **S4 absence pass** →
  **S5 citation gate** (every clause-based finding must quote the contract verbatim or it is
  dropped in Go).

Scoring is deterministic Go, never an LLM judging an LLM. Primary metric: recall over a
corpus of twelve synthetic contracts with hand-labelled ground truth. The headline number is
**absence-detection rate**, where the baseline sits near zero.

See [REPRODUCTION.md](REPRODUCTION.md) for how to run it, and how judges can verify the
committed numbers without an API key.

## What existed before, and what was built during the competition

Nothing existed before. Every line of code, every contract, and the entire legal corpus were
built during the micro1 Agentic Workflows Hackathon.

## Compliance

- Every contract in the corpus is **synthetic**, written from scratch. No real employment
  contract from the internet is used.
- No real personal data. All company and person names are fictional.
- No credentials in the repository. Keys are supplied through the environment at run time;
  `.env.example` lists only variable names.
- A qualified human remains the final decision maker by design — the memo triages and refers,
  it never signs.
