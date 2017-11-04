package p092

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	before, start := (*ListNode)(nil), head
	for i := 1; start != nil && i < m; i++ {
		before, start = start, start.Next
	}

	if start == nil {
		return head
	}

	prev, cur := before, start

	for i := m; cur != nil && i <= n; i++ {
		next := cur.Next
		cur.Next = prev
		prev, cur = cur, next
	}

	if start == head {
		head = prev
	} else {
		before.Next = prev
	}
	start.Next = cur

	return head
}
