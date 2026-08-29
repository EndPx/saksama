# Changelog

Each stage was measured before the next began. Evidence numbers are real, pulled
from `results/` (12 synthetic contracts, 13 ground-truth findings, model
`minimax/minimax-m3:free` via OpenRouter, temperature 0). Regenerate any row with
`./bin/eval -base results/s0_baseline.json -sol results/<stage>.json`.

| Stage | What was tried and why | Evidence | Decision or lesson |
|-------|------------------------|----------|--------------------|
| Corpus | Excluded two provisions from Constitutional Court ruling 168/PUU-XXI/2023 — severance as a mandatory floor, and termination requiring bipartite negotiation — because neither could be verified against primary text within the time available. | Corpus fixed at 14 provisions, all traceable to primary or consistent-secondary sources; 2 claims dropped. | Ground truth that cannot be traced to a primary source is not ground truth; it is an assumption wearing a fact's clothes. Kept deliberately out of scope. |
| Model selection | Needed an OpenRouter free model that follows JSON. Tried `nvidia/nemotron-3-ultra-550b` and `-lightning` (reasoning models) and `google/gemma-4`, `z-ai/glm-5.2`. | Nemotron: ~150–210 s/call, dumps chain-of-thought into `content`, truncates before the JSON — a ~630-call pipeline would take ~26 h. Gemma/GLM: HTTP 429 on the shared free pool. `minimax/minimax-m3:free`: ~2–20 s/call, clean JSON. | Chose minimax-m3. Lesson: a reasoning model is the wrong tool for high-volume structured extraction; the whole staged run then took ~29 min at \$0. Robust JSON extraction (last balanced block, tolerant of fences/prose) and in-body-429 retries were required to make free models usable. |
| S0 baseline | One free-form review call, then a separate normalisation call into JSON (the naive approach). | recall 15.4%, precision 2.4%, absence 16.7%, tier acc 50%, cross-section 0%, conf-A recall 9.1%; 2 correct of 82 reported; 14 false positives on clean contracts. | Baseline is genuinely weak and noisy: it over-reports (82 findings, 2 right) and almost never catches what is missing. This is the zero point. |
| S1 structured | Same whole-document review but output forced to constrained JSON over the 14 ids. | recall 46.2%, precision 10.9%, absence 16.7%; 6 correct of 55. | Structure roughly triples recall over baseline, but whole-document review still floods findings and barely improves absence. |
| S2 sections | Split the contract by article and review one section at a time, to fit long documents. | recall 23.1%, precision 11.1%, absence 16.7%; 3 correct of 27. | **Made things worse.** Per-section review halved S1 recall: a section read in isolation loses cross-article context (the c11 forfeiture trap needs Pasal 4+9+14 together) and the model has less to anchor on. Kept in the corpus and here as a recorded negative result rather than deleted. |
| S3 checklist | Stop letting the agent choose what to look at; run one targeted check per clause-based provision. | recall 53.8%, precision 35.0% (best of any stage), absence 0.0%; 7 correct of 20. | The checklist is the biggest precision win — forcing one provision at a time cuts noise sharply. But by design it never asks what is absent, so absence detection is zero. |
| S4 absence | Add an explicit pass: for each `tidak_ada_klausa` provision ask whether the contract satisfies it; expand PP35-13 into nine sub-checks. | recall 84.6%, precision 28.2%, absence 66.7%; 11 correct of 39. | This is where absence detection appears — from 0% to 66.7%. The gain over S3 comes almost entirely from provisions that are *missing*, which is the whole thesis. |
| S5 citation gate | Require every clause-based finding to quote the contract verbatim (≤200 chars), verified in Go; drop the rest. | recall 92.3%, precision 27.3%, absence 83.3%, tier acc 100%, cross-section 100%, conf-A recall 90.9%; 12 correct of 44. | Final solution. The gate removes ungrounded clause findings; net recall and absence both rise versus S4 (stages are independent LLM runs, so they are not perfectly monotonic). |

## Main failure mode and hot take

**Failure mode:** precision stays low (27% at S5) — the reviewer still over-reports,
especially borderline provisions on the clean and challenging contracts. Recall and
absence detection are strong; separating "plausible" from "correct" on the present
clauses is the remaining weakness.

**Hot take, with evidence:** an agent that is reactive to text systematically misses
what is absent from that text. The absence-detection rate is flat at ~17% from the
baseline all the way through S3 (structured output, per-section reading, even a
clause-by-clause checklist) — and only jumps to 66.7% and then 83.3% once an explicit
pass **forces the agent to ask what is missing** (S4, S5). Reliability came from a
checklist that asks "what should be here that isn't", not from a larger model or a
cleverer prompt.
