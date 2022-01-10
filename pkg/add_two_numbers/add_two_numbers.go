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
	var (
		currentVal, carry int
	)
	for l1 != nil || l2 != nil {
		var (
			valueL1, valueL2 int
		)
		if l1 != nil {
			valueL1 = l1.Val
		}
		if l2 != nil {
			valueL2 = l2.Val
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
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		// create a new node only when is necessary
		if l1 != nil || l2 != nil {
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
