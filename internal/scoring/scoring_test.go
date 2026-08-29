package scoring

import (
	"path/filepath"
	"testing"

	"github.com/EndPx/saksama/internal/statutes"
)

func loadAll(t *testing.T) (*statutes.Corpus, map[string]Truth) {
	t.Helper()
	corpus, err := statutes.Load(filepath.Join("..", "..", "data", "statutes", "2026-08.yaml"))
	if err != nil {
		t.Fatalf("load corpus: %v", err)
	}
	ids := []string{"c01", "c02", "c03", "c04", "c05", "c06", "c07", "c08", "c09", "c10", "c11", "c12"}
	truths := map[string]Truth{}
	for _, id := range ids {
		tr, err := LoadTruth(filepath.Join("..", "..", "data", "contracts", id, "truth.yaml"))
		if err != nil {
			t.Fatalf("load truth %s: %v", id, err)
		}
		truths[id] = tr
	}
	return corpus, truths
}

// perfectResult builds a stage result that reports exactly the ground truth,
// with the correct tier, so the scorer should return recall and precision 1.0.
func perfectResult(corpus *statutes.Corpus, truths map[string]Truth) StageResult {
	res := StageResult{Stage: "oracle"}
	for _, id := range []string{"c01", "c02", "c03", "c04", "c05", "c06", "c07", "c08", "c09", "c10", "c11", "c12"} {
		tr := truths[id]
		cr := ContractResult{ContractID: id}
		for _, tf := range tr.Findings {
			p, _ := corpus.Get(tf.StatuteID)
			cr.Findings = append(cr.Findings, Finding{
				StatuteID: tf.StatuteID,
				Section:   tf.Section,
				Tier:      p.Tier,
				Deteksi:   p.Deteksi,
			})
		}
		res.Contracts = append(res.Contracts, cr)
	}
	return res
}

func TestPerfectScoresOne(t *testing.T) {
	corpus, truths := loadAll(t)
	m := Evaluate(corpus, truths, perfectResult(corpus, truths))
	if m.Recall != 1.0 {
		t.Errorf("recall = %v, want 1.0", m.Recall)
	}
	if m.Precision != 1.0 {
		t.Errorf("precision = %v, want 1.0", m.Precision)
	}
	if m.AbsenceRate != 1.0 {
		t.Errorf("absence rate = %v, want 1.0", m.AbsenceRate)
	}
	if m.TierAccuracy != 1.0 {
		t.Errorf("tier accuracy = %v, want 1.0", m.TierAccuracy)
	}
	if m.CrossRecall != 1.0 {
		t.Errorf("cross recall = %v, want 1.0", m.CrossRecall)
	}
	if m.FalsePositivesClean != 0 {
		t.Errorf("false positives on clean = %d, want 0", m.FalsePositivesClean)
	}
	// Corpus totals: c01=2, c04=1, c05=4, c06=1, c07=1, c08=1, c11=1, c12=2 = 13.
	if m.TruthTotal != 13 {
		t.Errorf("truth total = %d, want 13", m.TruthTotal)
	}
	// Absence findings: c01 PP35-15 x1 + c05 x4 + c06 x1 = 6.
	if m.AbsenceTruth != 6 {
		t.Errorf("absence truth = %d, want 6", m.AbsenceTruth)
	}
}

func TestFalsePositivesCountedOnClean(t *testing.T) {
	corpus, truths := loadAll(t)
	res := StageResult{Stage: "noisy", Contracts: []ContractResult{
		{ContractID: "c02", Findings: []Finding{{StatuteID: "PP35-8", Section: "Pasal 1"}}},
		{ContractID: "c09", Findings: []Finding{{StatuteID: "SE-M5-2025", Section: "Pasal 4"}}},
	}}
	m := Evaluate(corpus, truths, res)
	if m.FalsePositivesClean != 2 {
		t.Errorf("false positives on clean = %d, want 2", m.FalsePositivesClean)
	}
	if m.Correct != 0 {
		t.Errorf("correct = %d, want 0 (clean contracts have no findings)", m.Correct)
	}
	if m.Precision != 0 {
		t.Errorf("precision = %v, want 0", m.Precision)
	}
}
