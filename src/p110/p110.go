package p110

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func treeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(treeHeight(root.Left), treeHeight(root.Right))
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	d := treeHeight(root.Left) - treeHeight(root.Right)
	if abs(d) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}
