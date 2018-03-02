package p145_1

func postorderTraversal(root *TreeNode) []int {
	vals := []int {}
	lefts := [](*TreeNode) {}

	for root != nil {
		vals = append([]int {root.Val}, vals...)
		if root.Left != nil {
			lefts = append([](*TreeNode) {root.Left}, lefts...)
		}
		root = root.Right
		if root == nil && len(lefts) != 0 {
			root, lefts = lefts[0], lefts[1:]
		}
	}

	return vals
}
