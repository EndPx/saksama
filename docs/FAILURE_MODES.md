# Failure modes (S5, final solution)

The final solution reports 18 findings; 13 are correct (recall 100% of the 13
ground-truth findings) and 5 are false positives (precision 72.2%). This document
makes those 5 auditable rather than hiding them. Machine-readable form:
[`results/failure_modes.json`](../results/failure_modes.json).

## The five false positives

| # | Contract | Rule | Article | Contract evidence | Why it fired | Why it is a false positive | Category |
|---|----------|------|---------|-------------------|--------------|----------------------------|----------|
| 1 | c01 | PP35-17 | Pasal 8 | "Uang kompensasi ... hanya akan diberikan ... apabila Perjanjian ini diperpanjang" | The clause conditions compensation — a genuine problem area. | The applicable rules (PP35-15/16) were already reported and matched; PP35-17 (early-termination compensation) is an adjacent, over-applied rule. | interpretation |
| 2 | c01 | PP35-13 | ABSENT | (all nine mandatory items are present in c01) | The nine-item sub-check judged a present item as missing. | c01 contains every mandatory Pasal 13 item. | extraction |
| 3 | c06 | PP35-26-31 | Pasal 5 | "jabatan PIHAK KEDUA tidak berhak atas upah kerja lembur" | The overtime-waiver clause is a real problem. | The waiver is governed by PP35-27-5 (correctly reported and matched); the 4h/18h cap and rates (PP35-26-31) are not the applicable rule. | interpretation |
| 4 | c11 | PP35-16 | Pasal 14 | "seluruh hak atas uang kompensasi ... menjadi gugur" | The forfeiture clause is genuinely problematic. | PP35-17 (the applicable rule) was reported and matched; PP35-16 (amount formula) is an adjacent over-applied rule. | interpretation |
| 5 | c11 | PP35-15 | ABSENT | (Pasal 9 grants compensation; Pasal 14 forfeits it) | The forfeiture nullifies compensation, so the absence pass read the protection as unsatisfied. | Debatable — compensation is present (Pasal 9); the real defect is the forfeiture, captured by PP35-17. Reading it as "effectively absent" is defensible. | genuinely ambiguous |

## Failure-mode summary

| Failure mode | Count | Category | Status |
|--------------|-------|----------|--------|
| Adjacent-rule over-application on a genuinely risky clause | 3 | interpretation error | Known limitation |
| Absence sub-check judged a present item as missing | 1 | extraction error | Known limitation |
| Forfeiture read as effectively-absent compensation | 1 | genuinely ambiguous | Borderline (defensible) |

## Key observation

Every remaining false positive lands on a clause that is **itself a real problem
area** (a conditional-compensation clause, an overtime waiver, a forfeiture clause).
The engine over-tags an *adjacent* rule; it does not hallucinate findings on
compliant text. This is confirmed by **zero false positives on the clean contracts**
(c02, c09, c10). For a triage tool this is a benign failure mode: it points the
reader at the right clause and occasionally cites one rule too many, rather than
inventing risk where none exists.

These five are **not automatically "fixed"**: three are adjacent-rule noise that
tightening prompts could reduce at some recall risk, one is an extraction slip, and
one is a legitimate interpretation question (is a forfeited protection "present but
void" or "effectively absent"?). They are recorded as known limitations rather than
silently patched.

## Citation-accuracy caveat

`results/audit.md` reports citation presence and location accuracy at 100% for S5.
That metric is **deterministic and syntactic/location-based, not semantic**: it
verifies the quoted excerpt exists in the contract and lies within the cited
article, but not that the excerpt *semantically supports* the finding. Semantic
citation verification would require an LLM judge and is intentionally out of scope
for the deterministic scorer; it is future work.
