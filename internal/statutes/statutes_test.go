package statutes

import (
	"path/filepath"
	"testing"
)

func corpusPath() string {
	return filepath.Join("..", "..", "data", "statutes", "2026-08.yaml")
}

func loadOrFatal(t *testing.T) *Corpus {
	t.Helper()
	c, err := Load(corpusPath())
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	return c
}

func TestCorpusLoadsFourteen(t *testing.T) {
	c := loadOrFatal(t)
	if got := len(c.Provisions); got != 14 {
		t.Fatalf("want 14 provisions, got %d", got)
	}
}

func TestIDsUnique(t *testing.T) {
	c := loadOrFatal(t)
	seen := map[string]bool{}
	for _, p := range c.Provisions {
		if seen[p.ID] {
			t.Errorf("duplicate id %q", p.ID)
		}
		seen[p.ID] = true
	}
}

func TestTiersValid(t *testing.T) {
	c := loadOrFatal(t)
	for _, p := range c.Provisions {
		if !validTiers[p.Tier] {
			t.Errorf("%s: invalid tier %q", p.ID, p.Tier)
		}
	}
}

func TestConfidencePopulated(t *testing.T) {
	c := loadOrFatal(t)
	for _, p := range c.Provisions {
		if p.Confidence != ConfidenceA && p.Confidence != ConfidenceB {
			t.Errorf("%s: confidence not populated (%q)", p.ID, p.Confidence)
		}
	}
}

func TestDetectionValid(t *testing.T) {
	c := loadOrFatal(t)
	for _, p := range c.Provisions {
		if !validDeteksi[p.Deteksi] {
			t.Errorf("%s: invalid deteksi %q", p.ID, p.Deteksi)
		}
	}
}

func TestExpectedIDsPresent(t *testing.T) {
	c := loadOrFatal(t)
	want := []string{
		"PP35-12", "UU13-60-1", "UU13-60-2", "PP35-4-2", "PP35-8", "PP35-13",
		"PP35-14", "PP35-15", "PP35-16", "PP35-17", "PP35-26-31", "PP35-27-5",
		"MK168-79-2b", "SE-M5-2025",
	}
	for _, id := range want {
		if _, ok := c.Get(id); !ok {
			t.Errorf("missing expected provision id %q", id)
		}
	}
}
