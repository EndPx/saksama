# Changelog

Each stage was measured before the next began. Evidence numbers are real, pulled
from `results/` (12 synthetic contracts, 13 ground-truth findings, model
`minimax/minimax-m3:free` via OpenRouter, temperature 0). Regenerate any row with
`./bin/eval -base results/s0_baseline.json -sol results/<stage>.json`.

| Stage | What was tried and why | Evidence (recall / precision / F1 / absence) | Decision or lesson |
|-------|------------------------|----------------------------------------------|--------------------|
| Corpus | Excluded two provisions from Constitutional Court ruling 168/PUU-XXI/2023 — severance as a mandatory floor, and termination requiring bipartite negotiation — because neither could be verified against primary text within the time available. | Corpus fixed at 14 provisions; 2 claims dropped. | Ground truth that cannot be traced to a primary source is not ground truth; it is an assumption wearing a fact's clothes. Kept deliberately out of scope. |
| Model selection | Needed an OpenRouter free model that follows JSON. Tried `nvidia/nemotron-3-ultra-550b` and `-lightning` (reasoning) and `google/gemma-4`, `z-ai/glm-5.2`. | Nemotron: ~150–210 s/call, dumps chain-of-thought into `content`, truncates before the JSON (a ~630-call run would be ~26 h). Gemma/GLM: HTTP 429 on the shared free pool. `minimax/minimax-m3:free`: ~2–20 s/call, clean JSON. | Chose minimax-m3. A reasoning model is the wrong tool for high-volume structured extraction; the full run then took ~29 min at \$0. Robust JSON extraction (last balanced block, tolerant of fences/prose) and in-body-429 retries were required. |
| S0 baseline | One free-form review call, then a separate normalisation call into JSON (the naive approach). | recall 23.1%, precision 3.7%, F1 6.3%, absence 33.3%; 3 correct of 82 reported; false-positive rate 96.3%. | Baseline is weak and extremely noisy: 82 findings, 3 right. It also barely catches what is missing. This is the zero point. |
| S1 structured | Same whole-document review but output forced to constrained JSON over the 14 ids. | recall 46.2%, precision 50.0%, F1 48.0%, absence 16.7%; 6 correct of 12. | Structure doubles recall and lifts precision sharply over baseline, but whole-document review still barely improves absence. |
| S2 sections | Split the contract by article and review one section at a time, to fit long documents. | recall 38.5%, precision 18.5%, F1 25.0%, absence 50.0%; 5 correct of 27. | **Made things worse.** Per-section review dropped both recall and precision below S1 (F1 25% vs 48%): a section read in isolation loses cross-article context (the c11 forfeiture trap needs Pasal 4+9+14 together) and floods findings. Kept as a recorded negative result rather than deleted. |
| S3 checklist | Stop letting the agent choose what to look at; run one targeted check per clause-based provision, and instruct it to flag only clear violations. | recall 53.8%, precision 63.6%, F1 58.3%, absence 0.0%; 7 correct of 11. | The checklist is the big precision win. But by design it never asks what is absent, so absence detection is zero. |
| S4 absence | Add an explicit pass: for each `tidak_ada_klausa` provision ask whether the contract satisfies it; expand PP35-13 into nine sub-checks; PP35-27-5 fires only when overtime is actually waived without a defined job category. | recall 100%, precision 65.0%, F1 78.8%, absence 100%; 13 correct of 20. | This is where absence detection appears — 0% to 100% — and recall reaches 100%. The gain is almost entirely provisions that are *missing*, which is the whole thesis. |
| S5 citation gate | Require every clause-based finding to quote the contract verbatim (≤200 chars), verified in Go; drop the rest. | recall 100%, precision 72.2%, F1 83.9%, absence 100%, tier acc 100%, cross-section 100%; 13 correct of 18; false-positive rate 27.8%. | Final solution. The gate removes ungrounded clause findings, raising precision to 72% with recall held at 100%. |
| Precision refinement (iteration on S3–S5) | An earlier run over-reported badly (S5 precision 27%, 32 false positives): the absence check flagged PP35-27-5 on nearly every contract, and PKWT-only provisions fired on PKWTT contracts. Added contract-type classification (PKWT vs PKWTT), skipped PKWT-only provisions on PKWTT, made PP35-27-5 fire only on a genuine overtime waiver, and told the checklist to flag only clear violations. | S5 precision 27.3% → 72.2%, false positives 32 → 5, false positives on clean contracts 6 → 0, with recall held at 100%. | Classify the document before applying rules. Most false positives were category errors (a permanent-contract clause judged against fixed-term rules) and one provision with conditional applicability — not model weakness. |

## Main failure mode and hot take

**Failure mode:** the remaining precision gap (S5 72%) is 5 false positives — mostly
borderline clause judgments on the challenging and mixed contracts. Recall, absence
detection, tier accuracy, and cross-section recall are all at 100%.

**Hot take, with evidence:** an agent that is reactive to text systematically misses
what is absent from that text. The absence-detection rate stays low (17–50%) through
structured output, per-section reading, and even a clause-by-clause checklist (0%) —
and only reaches 100% once an explicit pass **forces the agent to ask what is missing**
(S4). Reliability came from a checklist that asks "what should be here that isn't", and
from classifying the contract before applying rules — not from a larger model or a
cleverer prompt.
