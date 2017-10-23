package p061

func lengthOfList(head *ListNode) int {
	l := 0
	for ; head != nil; head = head.Next {
		l++
	}
	return l
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	ll := lengthOfList(head)
	if k >= ll {
		return rotateRight(head, k % ll)
	}

	prev, cur := head, head
	i := 0
	for ; i < k; i++ {
		cur = cur.Next
	}

	for ; cur.Next != nil; prev, cur = prev.Next, cur.Next {}

	cur.Next, head = head, prev.Next
	prev.Next = nil

	return head
}
