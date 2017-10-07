package p024

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	for pre, cur := head, head.Next; cur != nil; {
		*pre, *cur = *cur, *pre
		pre.Next, cur.Next = cur, pre.Next

		pre = cur.Next
		if pre != nil {
			cur = pre.Next
		} else {
			break
		}
	}

	return head
}
