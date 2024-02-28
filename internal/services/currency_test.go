package currency

import "testing"

func TestSplitCodeIntoPair(t *testing.T) {
	var base, target string
	base, target, err := SplitCodeIntoPair("EUR/USD")

	if base != "EUR" || target != "USD" || err != nil {
		t.Errorf("Expected EUR and USD, got %s and %s", base, target)
	}
}

func TestValidateCode(t *testing.T) {
	testCases := []struct {
		title    string
		code     string
		expected bool
	}{
		{
			title:    "valid code",
			code:     "MXN/EUR",
			expected: true,
		},
		{
			title:    "invalid code without slash",
			code:     "MXNEUR",
			expected: false,
		},
		{
			title:    "invalid code short len",
			code:     "MXN_EUR",
			expected: false,
		},
		{
			title:    "invalid code long len",
			code:     "MXNEUR/USDRUB",
			expected: false,
		},
		// {
		// 	title:    "code not in db",
		// 	code:     "QWE/RTY",
		// 	expected: false,
		// },
	}

	for _, tc := range testCases {
		res := validateCode(tc.code)
		if res != tc.expected {
			t.Errorf("Test failed on: %s", tc.title)
		}
	}
}
