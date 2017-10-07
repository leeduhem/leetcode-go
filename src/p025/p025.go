package p025

func reverseList(head, tail *ListNode) (head1*ListNode, tail1*ListNode) {
	tail1 = head

	if head == nil {
		head1 = head
		return
	}

	pre := head
	for cur, next := head.Next, (*ListNode)(nil); cur != tail; pre, cur = cur, next {
		next, cur.Next = cur.Next, pre
	}
	head1 = pre

	return
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	ohead := &ListNode{ 0, head }

	sprev, shead, stail := ohead, ohead.Next, ohead.Next

OuterLoop:
	for {
		for cnt := 0 ; cnt < k; cnt++ {
			if stail == nil {
				break OuterLoop
			}
			stail = stail.Next
		}

		h, t := reverseList(shead, stail)
		sprev.Next = h
		t.Next = stail

		sprev = t
		shead = stail
	}

	return ohead.Next
}
