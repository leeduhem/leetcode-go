package p117

func connect(root *TreeLinkNode) {
	if root == nil {
		return
	}

	pre := root

	for pre != nil {
		var temp, cur *TreeLinkNode

		temp = &TreeLinkNode{ Val: 0 }
		cur = temp

		for pre != nil {
			if pre.Left != nil {
				cur.Next = pre.Left
				cur = cur.Next
			}
			if pre.Right != nil {
				cur.Next = pre.Right
				cur = cur.Next
			}
			pre = pre.Next
		}
		pre = temp.Next
	}
}
