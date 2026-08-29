package scoring

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func readStage(t *testing.T, path string) StageResult {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	var r StageResult
	if err := json.Unmarshal(b, &r); err != nil {
		t.Fatalf("parse %s: %v", path, err)
	}
	return r
}

// TestS5AuditFrozen pins the audit numbers to the committed (frozen) S5 result,
// so any accidental change to the scoring/audit logic is caught. These are the
// submission-baseline numbers reported in results/audit.md.
func TestS5AuditFrozen(t *testing.T) {
	corpus, truths := loadAll(t)
	res := readStage(t, filepath.Join("..", "..", "results", "s5_gated.json"))

	cm := ConfusionMatrix(corpus, truths, res)
	if cm.TP != 10 || cm.FP != 5 || cm.FN != 0 || cm.TN != 153 {
		t.Errorf("confusion = %+v, want TP=10 FP=5 FN=0 TN=153", cm)
	}

	ids := []string{"c01", "c02", "c03", "c04", "c05", "c06", "c07", "c08", "c09", "c10", "c11", "c12"}
	text := map[string]string{}
	for _, id := range ids {
		b, err := os.ReadFile(filepath.Join("..", "..", "data", "contracts", id, "contract.md"))
		if err != nil {
			t.Fatalf("read contract %s: %v", id, err)
		}
		text[id] = string(b)
	}
	cs := Citation(text, res)
	if cs.ClauseFindings == 0 || cs.QuoteInContract != cs.ClauseFindings {
		t.Errorf("citation presence not 100%%: %+v", cs)
	}
	if cs.QuoteInSection != cs.ClauseFindings {
		t.Errorf("citation location not 100%%: %+v", cs)
	}
}
