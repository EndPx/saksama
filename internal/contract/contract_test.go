package contract

import (
	"os"
	"path/filepath"
	"testing"
)

func read(t *testing.T, id string) string {
	t.Helper()
	b, err := os.ReadFile(filepath.Join("..", "..", "data", "contracts", id, "contract.md"))
	if err != nil {
		t.Fatalf("read %s: %v", id, err)
	}
	return string(b)
}

func TestParseC11HasSixteenArticles(t *testing.T) {
	pre, secs := Parse(read(t, "c11"))
	if pre == "" {
		t.Error("preamble should carry the parties' identities, got empty")
	}
	if len(secs) != 16 {
		t.Fatalf("want 16 articles, got %d", len(secs))
	}
	if secs[13].Number != "14" {
		t.Errorf("14th article number = %q, want 14", secs[13].Number)
	}
	if got := secs[13].Title; got == "" {
		t.Error("Pasal 14 title should not be empty")
	}
	if secs[13].Label() != "Pasal 14" {
		t.Errorf("label = %q, want Pasal 14", secs[13].Label())
	}
}

func TestContainsQuoteNormalisesWhitespace(t *testing.T) {
	text := "Pekerja\n  wajib   menjalani masa   percobaan"
	if !ContainsQuote(text, "wajib menjalani masa percobaan") {
		t.Error("quote should match after whitespace normalisation")
	}
	if ContainsQuote(text, "tidak ada di kontrak") {
		t.Error("absent quote must not match")
	}
	if ContainsQuote(text, "") {
		t.Error("empty quote must not match")
	}
}
