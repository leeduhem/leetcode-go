package p129

func sumNumbers(root *TreeNode) int {
	var sum func(root *TreeNode, s int) int
	sum = func(root *TreeNode, s int) int {
		if root == nil {
			return 0
		}
		if root.Left == nil && root.Right == nil {
			return s*10 + root.Val
		}
		return sum(root.Left, s * 10 + root.Val) + sum(root.Right, s * 10 + root.Val)
	}

	return sum(root, 0)
}
