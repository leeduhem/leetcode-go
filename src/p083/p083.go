package p083

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	for prev, cur := head, head.Next; cur != nil; cur = cur.Next {
		if prev.Val == cur.Val {
			prev.Next = cur.Next
		} else {
			prev = cur
		}
	}

	return head
}
