package longest_substring

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	testcases := []struct {
		Name       string
		StringCase string
		Expected   int
	}{
		{"Longest: abc", "abcabcbb", 3},
		{"Longest: wke", "pwwkew", 3},
		{"Longest: d", "cdd", 2},
	}
	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			result := lengthOfLongestSubstring(tc.StringCase)
			if result != tc.Expected {
				t.Fatalf("Failed for %s\n", tc.Name)
			}
		})
	}
}
