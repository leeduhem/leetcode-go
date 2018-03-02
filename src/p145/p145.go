package p145

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int {}
	}

	left, right := postorderTraversal(root.Left), postorderTraversal(root.Right)
	return append(append(left, right...), root.Val)
}
