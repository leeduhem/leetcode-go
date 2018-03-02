package p144_1

func preorderTraversal(root *TreeNode) []int {
	vals := []int {}
	rights := [](*TreeNode) {}

	for root != nil {
		vals = append(vals, root.Val)
		if root.Right != nil {
			rights = append(rights, root.Right)
		}
		root = root.Left
		if root == nil && len(rights) != 0 {
			root, rights = rights[len(rights)-1], rights[:len(rights)-1]
		}
	}

	return vals
}
