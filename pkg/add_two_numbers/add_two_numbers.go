package add_two_numbers

// ListNode which holds a value and a pointer
// to the next node
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	tail := new(ListNode)
	iter := tail
	// this will keep l1 and l2 holding their current references intact
	auxL1, auxL2 := *l1, *l2
	pointerL1, pointerL2 := &auxL1, &auxL2
	var (
		currentVal, carry int
	)
	for pointerL1 != nil || pointerL2 != nil {
		var (
			valueL1, valueL2 int
		)
		if pointerL1 != nil {
			valueL1 = pointerL1.Val
		}
		if pointerL2 != nil {
			valueL2 = pointerL2.Val
		}

		currentVal = valueL1 + valueL2 + carry
		if currentVal > 9 {
			currentVal %= 10
			carry = 1
		} else {
			carry = 0
		}

		iter.Val = currentVal
		// reset values
		if pointerL1 != nil {
			pointerL1 = pointerL1.Next
		}
		if pointerL2 != nil {
			pointerL2 = pointerL2.Next
		}
		// create a new node only when is necessary
		if pointerL1 != nil || pointerL2 != nil {
			iter.Next = new(ListNode)
			iter = iter.Next
		}
	}
	if carry != 0 {
		iter.Next = &ListNode{
			Val: carry,
		}
	} else {
		iter = nil
	}
	return tail
}
