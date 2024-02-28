package currency

import "testing"

func TestSplitCodeIntoPair(t *testing.T) {
	var base, target string
	base, target = SplitCodeIntoPair("EUR/USD")

	if base != "EUR" || target != "USD" {
		t.Errorf("Expected EUR and USD, got %s and %s", base, target)
	}
}
