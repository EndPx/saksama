package scoring

import (
	"github.com/EndPx/saksama/internal/contract"
	"github.com/EndPx/saksama/internal/statutes"
)

// Confusion is a cell-level 2x2 confusion matrix over (contract x statute)
// decisions. This view exists to expose true negatives and make false positives
// auditable; it is a DIFFERENT granularity from the finding-level Metrics above:
// the nine PP35-13 sub-checks collapse into a single cell here. The finding-level
// precision/recall in Metrics is the headline; this matrix is the audit view.
type Confusion struct {
	TP, FP, FN, TN int
}

// Precision and Recall derived from the cell-level matrix.
func (c Confusion) Precision() float64 { return safeDiv(c.TP, c.TP+c.FP) }
func (c Confusion) Recall() float64    { return safeDiv(c.TP, c.TP+c.FN) }

// ConfusionMatrix computes the cell-level matrix. A cell (contract, statute) is
// positive in ground truth when that contract has >=1 finding for that statute,
// and positive in prediction when the stage reported >=1 finding for it.
func ConfusionMatrix(corpus *statutes.Corpus, truths map[string]Truth, res StageResult) Confusion {
	var cm Confusion
	for _, cr := range res.Contracts {
		t := truths[cr.ContractID]
		truthPos := map[string]bool{}
		for _, f := range t.Findings {
			truthPos[f.StatuteID] = true
		}
		predPos := map[string]bool{}
		for _, f := range cr.Findings {
			predPos[f.StatuteID] = true
		}
		for _, p := range corpus.Provisions {
			gt, pr := truthPos[p.ID], predPos[p.ID]
			switch {
			case gt && pr:
				cm.TP++
			case !gt && pr:
				cm.FP++
			case gt && !pr:
				cm.FN++
			default:
				cm.TN++
			}
		}
	}
	return cm
}

// CitationStats measures how well clause-based findings are grounded in the
// contract text.
//
// IMPORTANT — definition and its limits: this metric is DETERMINISTIC and
// SYNTACTIC/location-based, NOT semantic. It verifies that (a) a concrete
// verbatim excerpt exists (kutipan is non-empty), (b) that excerpt is present
// in the contract text after whitespace normalisation, and (c) that the excerpt
// lies within the article the finding cites. It does NOT verify that the
// excerpt semantically supports the finding — that judgement would require an
// LLM and is intentionally out of scope for deterministic scoring. Reported as
// PresenceAccuracy (b) and LocationAccuracy (c); semantic support is future work.
type CitationStats struct {
	ClauseFindings  int // findings with deteksi = ada_klausa
	WithQuote       int // of those, ones carrying a non-empty kutipan
	QuoteInContract int // of those, quote found verbatim in the contract
	QuoteInSection  int // of those, quote found within the cited article body
}

func (c CitationStats) PresenceAccuracy() float64 {
	return safeDiv(c.QuoteInContract, c.ClauseFindings)
}
func (c CitationStats) LocationAccuracy() float64 { return safeDiv(c.QuoteInSection, c.ClauseFindings) }

// Citation computes citation grounding over a stage result. contractsText maps
// contract id to the raw contract markdown.
func Citation(contractsText map[string]string, res StageResult) CitationStats {
	var cs CitationStats
	for _, cr := range res.Contracts {
		text := contractsText[cr.ContractID]
		_, sections := contract.Parse(text)
		byNum := make(map[string]string, len(sections))
		for _, s := range sections {
			byNum[s.Number] = s.Body
		}
		for _, f := range cr.Findings {
			if f.Deteksi != statutes.DeteksiAdaKlausa {
				continue
			}
			cs.ClauseFindings++
			if f.Kutipan == "" {
				continue
			}
			cs.WithQuote++
			if contract.ContainsQuote(text, f.Kutipan) {
				cs.QuoteInContract++
			}
			if body, ok := byNum[normSection(f.Section)]; ok && contract.ContainsQuote(body, f.Kutipan) {
				cs.QuoteInSection++
			}
		}
	}
	return cs
}
