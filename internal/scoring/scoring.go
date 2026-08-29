// Package scoring holds the shared result types and the deterministic metrics.
// No LLM is used here: matching a reported finding to a ground-truth finding is
// pure Go, so every number a stage produces is reproducible byte-for-byte.
package scoring

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/EndPx/saksama/internal/statutes"
	"gopkg.in/yaml.v3"
)

// Finding is one reported issue, from baseline normalisation or a solution stage.
type Finding struct {
	StatuteID string             `json:"statute_id"`
	Section   string             `json:"section"`
	Deskripsi string             `json:"deskripsi,omitempty"`
	Tier      statutes.Tier      `json:"tier,omitempty"`
	Deteksi   statutes.Detection `json:"deteksi,omitempty"`
	Kutipan   string             `json:"kutipan,omitempty"`
}

// TruthFinding is one ground-truth issue for a contract.
type TruthFinding struct {
	FindingID    string             `yaml:"finding_id" json:"finding_id"`
	StatuteID    string             `yaml:"statute_id" json:"statute_id"`
	Section      string             `yaml:"section" json:"section"`
	Tier         statutes.Tier      `yaml:"tier" json:"tier"`
	Deteksi      statutes.Detection `yaml:"deteksi" json:"deteksi"`
	CrossSection bool               `yaml:"cross_section" json:"cross_section"`
	Catatan      string             `yaml:"catatan" json:"catatan"`
}

// Truth is the ground truth for one contract.
type Truth struct {
	ContractID string         `yaml:"contract_id" json:"contract_id"`
	Jenis      string         `yaml:"jenis" json:"jenis"`
	SkalaUsaha string         `yaml:"skala_usaha" json:"skala_usaha"`
	Findings   []TruthFinding `yaml:"findings" json:"findings"`
}

// LoadTruth reads and parses a truth.yaml file.
func LoadTruth(path string) (Truth, error) {
	var t Truth
	data, err := os.ReadFile(path)
	if err != nil {
		return t, fmt.Errorf("read truth: %w", err)
	}
	if err := yaml.Unmarshal(data, &t); err != nil {
		return t, fmt.Errorf("parse truth: %w", err)
	}
	return t, nil
}

// Usage records token usage and cost for an LLM interaction.
type Usage struct {
	InputTokens  int     `json:"input_tokens"`
	OutputTokens int     `json:"output_tokens"`
	CostUSD      float64 `json:"cost_usd"`
}

// Add accumulates another usage into u.
func (u *Usage) Add(o Usage) {
	u.InputTokens += o.InputTokens
	u.OutputTokens += o.OutputTokens
	u.CostUSD += o.CostUSD
}

// ContractResult is a stage's reported findings for one contract.
type ContractResult struct {
	ContractID string    `json:"contract_id"`
	Findings   []Finding `json:"findings"`
	Usage      Usage     `json:"usage"`
	DurationMs int64     `json:"duration_ms"`
}

// StageResult is the full output of one stage across the corpus. It is the
// on-disk shape of results/sN_*.json.
type StageResult struct {
	Stage     string           `json:"stage"`
	Model     string           `json:"model"`
	Contracts []ContractResult `json:"contracts"`
	// NormalizationUsage is tracked separately (baseline fairness) and never
	// counted as part of the solution cost.
	NormalizationUsage *Usage `json:"normalization_usage,omitempty"`
}

// cleanContracts are the contracts with no ground-truth findings; any finding
// reported against them is a false positive.
var cleanContracts = map[string]bool{"c02": true, "c09": true, "c10": true}

var digits = regexp.MustCompile(`\d+`)

// normSection reduces a section label to a comparable key: "ABSENT" stays
// "ABSENT"; otherwise the first run of digits ("Pasal 14" -> "14"); failing
// that, the lower-cased trimmed string.
func normSection(s string) string {
	t := strings.TrimSpace(s)
	if strings.EqualFold(t, "ABSENT") {
		return "ABSENT"
	}
	if m := digits.FindString(t); m != "" {
		return m
	}
	return strings.ToLower(t)
}

// detOf returns the authoritative detection mode for a statute id, preferring
// the corpus and falling back to the ground-truth copy.
func detOf(corpus *statutes.Corpus, id string, fallback statutes.Detection) statutes.Detection {
	if p, ok := corpus.Get(id); ok {
		return p.Deteksi
	}
	return fallback
}

// Metrics is the deterministic scorecard for one stage over the corpus.
type Metrics struct {
	Stage string `json:"stage"`

	TruthTotal    int `json:"truth_total"`
	ReportedTotal int `json:"reported_total"`
	Correct       int `json:"correct"`

	Recall    float64 `json:"recall"`
	Precision float64 `json:"precision"`

	AbsenceTruth int     `json:"absence_truth"`
	AbsenceHit   int     `json:"absence_hit"`
	AbsenceRate  float64 `json:"absence_rate"`

	TierCorrect  int     `json:"tier_correct"`
	TierAccuracy float64 `json:"tier_accuracy"`

	CrossTruth  int     `json:"cross_truth"`
	CrossHit    int     `json:"cross_hit"`
	CrossRecall float64 `json:"cross_recall"`

	FalsePositivesClean int `json:"false_positives_clean"`

	ConfATruth  int     `json:"conf_a_truth"`
	ConfAHit    int     `json:"conf_a_hit"`
	ConfARecall float64 `json:"conf_a_recall"`

	TotalCostUSD     float64 `json:"total_cost_usd"`
	CostPerContract  float64 `json:"cost_per_contract"`
	AvgDurationMs    float64 `json:"avg_duration_ms"`
	NormalizationUSD float64 `json:"normalization_usd,omitempty"`
}

func safeDiv(a, b int) float64 {
	if b == 0 {
		return 0
	}
	return float64(a) / float64(b)
}

// Evaluate scores a stage result against the ground truth. truths is keyed by
// contract id. Matching is per contract: each ground-truth finding may be
// matched at most once, and each reported finding used at most once.
func Evaluate(corpus *statutes.Corpus, truths map[string]Truth, res StageResult) Metrics {
	m := Metrics{Stage: res.Stage}

	confA := func(id string) bool {
		p, ok := corpus.Get(id)
		return ok && p.Confidence == statutes.ConfidenceA
	}

	var totalDur int64
	for _, cr := range res.Contracts {
		truth := truths[cr.ContractID]
		m.TruthTotal += len(truth.Findings)
		m.ReportedTotal += len(cr.Findings)
		m.TotalCostUSD += cr.Usage.CostUSD
		totalDur += cr.DurationMs

		if cleanContracts[cr.ContractID] {
			m.FalsePositivesClean += len(cr.Findings)
		}

		used := make([]bool, len(cr.Findings))
		for _, tf := range truth.Findings {
			det := detOf(corpus, tf.StatuteID, tf.Deteksi)
			isAbsence := det == statutes.DeteksiTidakAdaKlausa
			if isAbsence {
				m.AbsenceTruth++
			}
			if tf.CrossSection {
				m.CrossTruth++
			}
			if confA(tf.StatuteID) {
				m.ConfATruth++
			}

			// Find an unused reported finding that matches.
			hit := -1
			for i, rf := range cr.Findings {
				if used[i] || rf.StatuteID != tf.StatuteID {
					continue
				}
				if !isAbsence && normSection(rf.Section) != normSection(tf.Section) {
					continue
				}
				hit = i
				break
			}
			if hit < 0 {
				continue
			}
			used[hit] = true
			m.Correct++
			if isAbsence {
				m.AbsenceHit++
			}
			if tf.CrossSection {
				m.CrossHit++
			}
			if confA(tf.StatuteID) {
				m.ConfAHit++
			}
			// Tier accuracy: reported tier equals the statute's tier.
			if p, ok := corpus.Get(tf.StatuteID); ok && cr.Findings[hit].Tier == p.Tier {
				m.TierCorrect++
			}
		}
	}

	m.Recall = safeDiv(m.Correct, m.TruthTotal)
	m.Precision = safeDiv(m.Correct, m.ReportedTotal)
	m.AbsenceRate = safeDiv(m.AbsenceHit, m.AbsenceTruth)
	m.TierAccuracy = safeDiv(m.TierCorrect, m.Correct)
	m.CrossRecall = safeDiv(m.CrossHit, m.CrossTruth)
	m.ConfARecall = safeDiv(m.ConfAHit, m.ConfATruth)
	if n := len(res.Contracts); n > 0 {
		m.CostPerContract = m.TotalCostUSD / float64(n)
		m.AvgDurationMs = float64(totalDur) / float64(n)
	}
	if res.NormalizationUsage != nil {
		m.NormalizationUSD = res.NormalizationUsage.CostUSD
	}
	return m
}
