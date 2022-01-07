package longest_palindromic_substring

import "testing"

func TestLongestPalindrome(t *testing.T) {
	// TODO: add golden files for larger inputs
	testcases := []struct {
		Name       string
		StringCase string
		Expected   string
	}{
		{"Longest Palindrome of babad", "babad", "bab"},
		{"Longest Palindrome of cbbd", "cbbd", "bb"},
		{"Longest Palindrome of aacabdkacaa", "aacabdkacaa", "aca"},
		{"Longest Palindrome of abcda", "abcda", "a"},
		{"Longest Palindrome of aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
	}
	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			result := longestPalindrome(tc.StringCase)
			if result != tc.Expected {
				t.Fatalf("Failed for %s\n", tc.Name)
			}
		})
	}
}
