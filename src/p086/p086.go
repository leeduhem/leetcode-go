package p086

func partition(head *ListNode, x int) *ListNode {
	var head1, head2 *ListNode
	tail1, tail2 := &head1, &head2

	for ; head != nil; head = head.Next {
		if head.Val < x {
			*tail1 = head
			tail1 = &(*tail1).Next
		} else {
			*tail2 = head
			tail2 = &(*tail2).Next
		}
	}

	if head1 != nil {
		*tail1 = head2
	}
	if *tail2 != nil {
		*tail2 = nil
	}

	if head1 != nil {
		return head1
	}
	return head2
}
