package p203

func removeElements(head *ListNode, val int) *ListNode {
	head1 := &ListNode{Next: head}
	for pre, cur := head1, head1.Next; cur != nil; {
		if cur.Val == val {
			cur = cur.Next
			pre.Next = cur
		} else {
			pre, cur = cur, cur.Next
		}
	}
	return head1.Next
}
