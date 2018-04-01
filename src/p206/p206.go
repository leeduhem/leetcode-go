package p206

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	head1 := reverseList(head.Next)
	head.Next = nil

	tail := head1;
	for ; tail.Next != nil; tail = tail.Next {}
	tail.Next = head

	return head1
}
