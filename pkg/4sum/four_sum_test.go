package four_sum

import (
	"reflect"
	"testing"
)

func TestFourSum(t *testing.T) {
	testcases := []struct {
		name     string
		nums     []int
		target   int
		expected [][]int
	}{
		{
			name:     "easy case",
			nums:     []int{1, 0, -1, 0, -2, 2},
			target:   0,
			expected: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
		{
			name:     "sorted case with double 0's",
			nums:     []int{-3, -2, -1, 0, 0, 1, 2, 3},
			target:   0,
			expected: [][]int{{-3, -2, 2, 3}, {-3, -1, 1, 3}, {-3, 0, 0, 3}, {-3, 0, 1, 2}, {-2, -1, 0, 3}, {-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result := fourSum(tc.nums, tc.target)

			check := reflect.DeepEqual(result, tc.expected)
			if !check {
				t.Errorf("invalid result:\ngot:\n%v\nexpected:\n%v\n", result, tc.expected)
			}
		})
	}
}
