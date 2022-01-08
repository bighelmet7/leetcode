package add_two_numbers

import (
	"testing"
)

func TestAddTwoNumbersTwoArrays(t *testing.T) {
	testcases := []struct {
		Name     string
		A        *ListNode
		B        *ListNode
		Expected []int
	}{
		{
			Name:     "[2,4,3] + [5,6,4] = [7,0,8]",
			A:        &ListNode{2, &ListNode{4, &ListNode{3, nil}}},
			B:        &ListNode{5, &ListNode{6, &ListNode{4, nil}}},
			Expected: []int{7, 0, 8},
		},
		{
			Name:     "[0] + [0] = [0]",
			A:        &ListNode{0, nil},
			B:        &ListNode{0, nil},
			Expected: []int{0},
		},
		{
			Name:     "[9,9,9] + [9] = [8,0,0,1]",
			A:        &ListNode{9, &ListNode{9, &ListNode{9, nil}}},
			B:        &ListNode{9, nil},
			Expected: []int{8, 0, 0, 1},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			result := addTwoNumbers(tc.A, tc.B)
			var iter int
			for result != nil {
				if result.Val != tc.Expected[iter] {
					t.Fatalf("Failed for %s\n", tc.Name)
				}
				result = result.Next
				iter++
			}
		})
	}
}

func TestAddTwoNumbersTwoArraysPassedValuesNotModified(t *testing.T) {
	A := &ListNode{1, nil}
	B := &ListNode{2, nil}
	expectedPointerA := A
	expectedPointerB := B

	// After calling the function both ListNode should still point
	// to the same address memory
	addTwoNumbers(A, B)
	if expectedPointerA != A {
		t.Fatalf("New address for A found. Expected: %p, Result: %p\n", expectedPointerA, A)
	}
	if expectedPointerB != B {
		t.Fatalf("New address for B found. Expected: %p, Result: %p\n", expectedPointerB, B)
	}
}
