package merge_two_sorted_lists

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	} else if list1 != nil && list2 == nil {
		return list1
	} else if list1 == nil && list2 != nil {
		return list2
	}

	aux1 := list1
	aux2 := list2
	head := &ListNode{}

	it := head
    if aux1.Val < aux2.Val {
        it.Val = aux1.Val
        aux1 = aux1.Next
    } else {
        it.Val = aux2.Val
        aux2 = aux2.Next
    }
    it.Next = &ListNode{}
    it = it.Next

	for aux1 != nil || aux2 != nil {
		if aux1 != nil && aux2 != nil {
			if aux1.Val < aux2.Val {
				it.Val = aux1.Val
				aux1 = aux1.Next
			} else {
				it.Val = aux2.Val
				aux2 = aux2.Next
			}
		} else if aux1 != nil && aux2 == nil {
			it.Val = aux1.Val
			aux1 = aux1.Next
		} else if aux1 == nil && aux2 != nil {
			it.Val = aux2.Val
			aux2 = aux2.Next
		}

		if aux1 != nil || aux2 != nil {
			it.Next = &ListNode{}
			it = it.Next
		}
	}

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
