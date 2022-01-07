package string_to_int

import "testing"

func TestMyAtoi(t *testing.T) {
	// Parametrized test
	testcases := []struct {
		Name       string
		StringCase string
		Expected   int
	}{
		{"Whitespaces with a negative", "    -42", -42},
		{"Numbers with letters", "1234 with words", 1234},
		{"Invalid sign", "+-12", 0},
		{"Min Int overflow", "-91283472332", -2147483648},
		{"Number", "1234", 1234},
	}
	// t.Run() it will execute test in parallel
	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			result := myAtoi(tc.StringCase)
			if result != tc.Expected {
				t.Fatalf("Failed for %s\n", tc.Name)
			}
		})
	}
}
