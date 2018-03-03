package p147

func insertionSortList(head *ListNode) *ListNode {
	new_head := &ListNode{ }

	for node, next := head, (*ListNode)(nil); node != nil; node = next {
		next = node.Next

		for prev := new_head; prev != nil; prev = prev.Next {
			if prev.Next == nil || prev.Next.Val > node.Val {
				prev.Next, node.Next = node, prev.Next
				break
			}
		}
	}

	return new_head.Next
}
