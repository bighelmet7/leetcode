package zigzag_conversion

import "testing"

func TestConvert(t *testing.T) {
	testcases := []struct {
		Name       string
		StringCase string
		NumOfRows  int
		Expected   string
	}{
		{"PAYPALISHIRING, rows: 3", "PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING, rows: 4", "PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"AB, rows: 1", "AB", 1, "AB"},
	}
	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			result := convert(tc.StringCase, tc.NumOfRows)
			if result != tc.Expected {
				t.Fatalf("Failed for %s\n", tc.Name)
			}
		})
	}
}
