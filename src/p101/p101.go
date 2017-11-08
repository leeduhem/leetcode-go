package p101

func isSymmetric(root *TreeNode) bool {
	var is_symmetric func (left, right *TreeNode) bool
	is_symmetric = func (left, right *TreeNode) bool {
		if left == nil || right == nil {
			return left == right
		}

		if left.Val != right.Val {
			return false
		}

		return is_symmetric(left.Left, right.Right) && is_symmetric(left.Right, right.Left)
	}

	return root == nil || is_symmetric(root.Left, root.Right)
}
