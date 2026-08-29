# Changelog

Each stage is measured before the next one begins. The Evidence column holds real numbers
pulled from `results/`, not qualitative description. Rows are filled in as stages land.

| Stage | What was tried and why | Evidence | Decision or lesson |
|-------|------------------------|----------|--------------------|
| Corpus | Excluded two provisions from Constitutional Court ruling 168/PUU-XXI/2023 — severance as a mandatory floor, and termination requiring bipartite negotiation — because neither could be verified against primary text within the available time. | 2 claims dropped; corpus narrowed to 14 provisions, all traceable to primary or consistent-secondary sources. | Ground truth that cannot be traced to a primary source is not ground truth; it is an assumption wearing a fact's clothes. Kept deliberately out of scope. |
| S0 baseline | _pending_ | _pending_ | _pending_ |
| S1 structured | _pending_ | _pending_ | _pending_ |
| S2 sections | _pending_ | _pending_ | _pending_ |
| S3 checklist | _pending_ | _pending_ | _pending_ |
| S4 absence | _pending_ | _pending_ | _pending_ |
| S5 citation gate | _pending_ | _pending_ | _pending_ |

## Main failure mode and hot take

_Hot take (to be evidenced by the baseline-vs-solution absence-detection rate):_ an agent that
is reactive to text will systematically miss what is absent from that text. Reliability comes
from a checklist that forces the agent to ask what is missing — not from a larger model or a
cleverer prompt.
