# Evaluation audit

Cell-level confusion matrix over (contract x statute) decisions. This exposes
true negatives; it is a coarser granularity than the finding-level headline
metrics (the nine PP35-13 sub-checks collapse to one cell here).

## Confusion matrix

### Baseline (S0)

```
                 Ground Truth
              Positive   Negative
Pred Positive      8         73
Pred Negative      2         85
```

Cell-level precision 9.9%, recall 80.0% (over 168 contract x statute cells).

### Solution (S5)

```
                 Ground Truth
              Positive   Negative
Pred Positive     10          5
Pred Negative      0        153
```

Cell-level precision 66.7%, recall 100.0% (over 168 contract x statute cells).

## Citation grounding

Deterministic and syntactic/location-based, NOT semantic (see internal/scoring/audit.go).

| Metric | Baseline (S0) | Solution (S5) |
|---|---|---|
| Clause-based findings | 32 | 10 |
| With a non-empty quote | 25 | 10 |
| Quote present in contract (presence accuracy) | 18.8% | 100.0% |
| Quote within cited article (location accuracy) | 9.4% | 100.0% |

Semantic support of the quote for the finding is not verified deterministically; see `docs/FAILURE_MODES.md`.
