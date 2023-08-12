package remove_the_nth_node

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n == 1 && head.Next == nil {
		return nil
	}

	l := []int{}
	for aux := head; aux != nil; aux = aux.Next {
		l = append(l, aux.Val)
	}

	var (
		prev    *ListNode
		current *ListNode
	)

	j := 1
	for i := len(l) - 1; i >= 0; i-- {
		if j != n {
			if current == nil {
				prev = &ListNode{Val: l[i]}
				current = prev
			} else {
				current = &ListNode{Val: l[i], Next: prev}
				prev = current
			}
		}
		j++
	}

	head = current
	return head
}

func Slice(head *ListNode) []int {
	n := head
	result := []int{}
	for n != nil {
		result = append(result, n.Val)
		n = n.Next
	}
	return result
}
