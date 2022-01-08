package palindrome_number

import "testing"

func TestIsPalindrome(t *testing.T) {
	testcases := []struct {
		Name        string
		IntegerCase int
		Expected    bool
	}{
		{"Palindrome: x = 121", 121, true},
		{"Not Palindrome: x = -121", -121, false},
		{"Not Palindrome: x = 10", 10, false},
		{"Not Palindrome: x = 123", 123, false},
		{"Palindrome: x = 1111111", 1111111, true},
		{"Palindrome: x = 0", 0, true},
	}
	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			result := isPalindrome(tc.IntegerCase)
			if result != tc.Expected {
				t.Fatalf("Failed for %s\n", tc.Name)
			}
		})
	}
}
