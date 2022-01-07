package median_sorted_arrays

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	testcases := []struct {
		Name     string
		ArrayA   []int
		ArrayB   []int
		Expected float64
	}{
		{"Median of an odd array", []int{1, 3}, []int{2}, 2.00000},
		{"Median of an even sorted array", []int{1, 2}, []int{3, 4}, 2.50000},
		{"Median of one empty and not empty array", []int{}, []int{2, 3}, 2.50000},
		{"Median of one empty and sorted array", []int{}, []int{1, 2, 3, 4, 5}, 3.00000},
	}
	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			result := findMedianSortedArrays(tc.ArrayA, tc.ArrayB)
			if result != tc.Expected {
				t.Fatalf("Failed for %s\n", tc.Name)
			}
		})
	}
}
