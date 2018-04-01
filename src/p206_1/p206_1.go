package p206_1

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for cur := head; cur != nil; {
		next := cur.Next
		cur.Next = pre
		pre, cur = cur, next
	}
	return pre
}
