package llm

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadEnvFile(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, ".env")
	content := "# a comment\n\nSAKSAMA_TEST_X=hello\nexport SAKSAMA_TEST_Z=\"quoted val\"\nSAKSAMA_TEST_Y=file\n"
	if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	// A var already in the environment must NOT be overridden by the file.
	t.Setenv("SAKSAMA_TEST_Y", "shell")
	os.Unsetenv("SAKSAMA_TEST_X")
	os.Unsetenv("SAKSAMA_TEST_Z")
	defer os.Unsetenv("SAKSAMA_TEST_X")
	defer os.Unsetenv("SAKSAMA_TEST_Z")

	if err := LoadEnvFile(p); err != nil {
		t.Fatalf("LoadEnvFile: %v", err)
	}
	if got := os.Getenv("SAKSAMA_TEST_X"); got != "hello" {
		t.Errorf("unset var not loaded: got %q, want hello", got)
	}
	if got := os.Getenv("SAKSAMA_TEST_Z"); got != "quoted val" {
		t.Errorf("quoted/export value not parsed: got %q, want 'quoted val'", got)
	}
	if got := os.Getenv("SAKSAMA_TEST_Y"); got != "shell" {
		t.Errorf("file overrode shell env: got %q, want shell", got)
	}
	// Missing file is not an error.
	if err := LoadEnvFile(filepath.Join(dir, "does-not-exist")); err != nil {
		t.Errorf("missing file should be nil, got %v", err)
	}
}

func TestCostUSD(t *testing.T) {
	c := &Client{priceIn: 3.0, priceOut: 15.0} // USD per million tokens
	got := c.costUSD(1_000_000, 1_000_000)
	if want := 18.0; got != want {
		t.Errorf("costUSD = %v, want %v", got, want)
	}
	if z := c.costUSD(0, 0); z != 0 {
		t.Errorf("zero usage cost = %v, want 0", z)
	}
}

func TestNewRequiresEnv(t *testing.T) {
	t.Setenv("SAKSAMA_API_BASE", "")
	t.Setenv("SAKSAMA_API_KEY", "")
	t.Setenv("SAKSAMA_MODEL", "")
	if _, err := New(); err == nil {
		t.Fatal("New() with empty env should error")
	}
}
