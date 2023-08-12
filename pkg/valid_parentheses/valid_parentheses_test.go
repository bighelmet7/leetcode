package valid_parentheses

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	testcases := []struct {
		name   string
		have   string
		expected bool
	}{
		{
			name:   "easy case",
			have:   "()",
			expected: true,
		},
		{
			name:   "mix",
			have:   "()[]{}",
			expected: true,
		},
		{
			name:   "invalid",
			have:   "(]",
			expected: false,
		},
		{
			name:   "closing order does not match",
			have:   "([)]",
			expected: false,
		},
		{
			name:   "mix valid closing",
			have:   "[{}]",
			expected: true,
		},
		{
			name:   "close first",
			have:   "){",
			expected: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result := isValid(tc.have)
			if result != tc.expected {
				t.Errorf("invalid result:\ngot:\n%v\nexpected:\n%v\n", result, tc.expected)
			}
		})
	}
}
