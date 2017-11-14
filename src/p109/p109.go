package p109

func listLength(head *ListNode) int {
	l := 0

	if head == nil {
		return l
	}

	for ; head != nil; head = head.Next {
		l++
	}

	return l
}

func sortedListToBST(head *ListNode) *TreeNode {
	l := listLength(head)

	if l == 0 {
		return nil
	}

	if l == 1 {
		return &TreeNode{ head.Val, nil, nil }
	}

	mid := l / 2
	prev := head
	for i := 0; i < mid-1; i++ {
		prev = prev.Next
	}

	val := prev.Next.Val
	right := sortedListToBST(prev.Next.Next)

	prev.Next = nil
	left := sortedListToBST(head)

	return &TreeNode{ val, left, right }
}
