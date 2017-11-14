package p114

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	prev := root.Left
	for ; prev != nil && prev.Right != nil; prev = prev.Right {}

	if prev != nil {
		root.Right, prev.Right = root.Left, root.Right
		root.Left = nil
	}

	return
}
