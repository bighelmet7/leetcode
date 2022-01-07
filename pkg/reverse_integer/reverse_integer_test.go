package reverse_integer

import "testing"

func TestReverse(t *testing.T) {
	testcases := []struct {
		Name     string
		Integer  int
		Expected int
	}{
		{"Integer: 123", 123, 321},
		{"Integer: -123", -123, -321},
		{"Integer: -2147483412", -2147483412, -2143847412},
		{"Integer: 1534236469", 1534236469, 0},
	}
	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			result := reverse(tc.Integer)
			if result != tc.Expected {
				t.Fatalf("Failed for %s\n", tc.Name)
			}
		})
	}
}
