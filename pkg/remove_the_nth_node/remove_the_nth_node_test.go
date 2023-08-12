package remove_the_nth_node

import (
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	testcases := []struct {
		name     string
		have     ListNode
		nth      int
		expected *ListNode
	}{
		{
			name:     "easy case",
			have:     ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			nth:      2,
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}},
		},
		{
			name:     "one node",
			have:     ListNode{Val: 5},
			nth:      1,
			expected: nil,
		},
		{
			name:     "remove tail",
            have:     ListNode{Val:3 , Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
			nth:      1,
            expected: &ListNode{Val: 3, Next: &ListNode{Val: 4}},
		},
		{
			name:     "two nodes",
            have:     ListNode{Val:3 , Next: &ListNode{Val: 4}},
			nth:      1,
            expected: &ListNode{Val: 3},
		},
		{
			name:     "remove head",
			have:     ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
			nth:      3,
			expected: &ListNode{Val: 4, Next: &ListNode{Val: 5}},
		},
		{
			name:     "123",
			have:     ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
			nth:      2,
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 3}},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			resultNodes := removeNthFromEnd(&tc.have, tc.nth)
			result := Slice(resultNodes)
			expected := Slice(tc.expected)

			check := reflect.DeepEqual(result, expected)
			if !check {
				t.Errorf("invalid result:\ngot:\n%v\nexpected:\n%v\n", result, expected)
			}
		})
	}
}
