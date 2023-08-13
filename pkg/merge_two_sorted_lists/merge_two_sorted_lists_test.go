package merge_two_sorted_lists

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	testcases := []struct {
		name     string
		have     []*ListNode
		expected *ListNode
	}{
		{
			name: "easy case",
			have: []*ListNode{
				{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}},
				{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
			},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4}}}}}},
		},
		{
			name:     "empty",
			have:     []*ListNode{nil, nil},
			expected: nil,
		},
		{
			name: "one node",
			have: []*ListNode{
				nil,
				{Val: 1},
			},
			expected: &ListNode{Val: 1},
		},
		{
			name: "different sizes",
			have: []*ListNode{
				{Val: 1, Next: &ListNode{Val: 2}},
				{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
			},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			resultNodes := mergeTwoLists(tc.have[0], tc.have[1])
			result := Slice(resultNodes)
			expected := Slice(tc.expected)

			check := reflect.DeepEqual(result, expected)
			if !check {
				t.Errorf("invalid result:\ngot:\n%v\nexpected:\n%v\n", result, expected)
			}
		})
	}
}
