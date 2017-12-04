package p124

import (
	"math"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32

	var maxPathDown func(root *TreeNode) int
	maxPathDown = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := max(0, maxPathDown(node.Left))
		right := max(0, maxPathDown(node.Right))
		maxSum = max(maxSum, left + right + node.Val)
		return max(left, right) + node.Val
	}

	maxPathDown(root)
	return maxSum
}
