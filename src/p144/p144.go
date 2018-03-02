package p144

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int {}
	}

	left, right := preorderTraversal(root.Left), preorderTraversal(root.Right)
	return append([]int { root.Val }, append(left, right...)...)
}
