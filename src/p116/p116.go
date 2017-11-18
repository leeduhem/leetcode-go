package p116

func connect(root *TreeLinkNode) {
	if root == nil {
		return
	}

	var pre, cur *TreeLinkNode

	pre = root
	for pre.Left != nil {
		cur = pre
		for cur != nil {
			cur.Left.Next = cur.Right
			if cur.Next != nil {
				cur.Right.Next = cur.Next.Left
			}
			cur = cur.Next
		}
		pre = pre.Left
	}
}
