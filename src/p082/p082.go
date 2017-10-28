package p082

func deleteDuplicates(head *ListNode) *ListNode {
	for cur := head; cur != nil; {
		for ; cur != nil && cur.Next != nil && cur.Val == cur.Next.Val; cur = cur.Next {}

		if head.Next == nil {
			return head
		}

		if head.Val == head.Next.Val {
			if cur == nil {
				head = nil
				return head
			}
			head = cur.Next
			cur  = head
			continue
		}

		break
	}

	if head == nil {
		return head
	}

	for prev, cur := head, head.Next; cur != nil; cur = cur.Next {
		for ; cur != nil && cur.Next != nil && cur.Val == cur.Next.Val; cur = cur.Next {}

		if cur == nil {
			prev.Next = nil
			return head
		}

		if prev.Next != cur {
			prev.Next = cur.Next
		} else {
			prev = cur
		}
	}

	return head
}
