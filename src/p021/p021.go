package p021

type ListNode struct {
    Val int
    Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{0, nil}
	end  := head

	for ; l1 != nil && l2 != nil; {
		if l1.Val < l2.Val {
			end.Next = l1
			l1 = l1.Next
		} else {
			end.Next = l2
			l2 = l2.Next
		}

		end = end.Next
	}

	if l1 != nil {
		end.Next = l1
	}

	if l2 != nil {
		end.Next = l2
	}

	return head.Next
}
