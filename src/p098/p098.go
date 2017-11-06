package p098

import (
	"math"
)

func isValid(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	if root.Val >= max || root.Val <= min {
		return false
	}

	return isValid(root.Left, min, root.Val) && isValid(root.Right, root.Val, max)
}

func isValidBST(root *TreeNode) bool {
	return isValid(root, math.MinInt64, math.MaxInt64)
}
