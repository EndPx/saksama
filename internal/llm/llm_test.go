package llm

import "testing"

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
