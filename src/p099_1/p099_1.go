package p099_1

import (
	"math"
)

func recoverTree(root *TreeNode) {
	var first, second, prev *TreeNode

	var traverse func (root *TreeNode)
	traverse = func (root *TreeNode) {
		if root == nil {
			return
		}

		traverse(root.Left)

		if first == nil && prev.Val >= root.Val {
			first = prev
		}

		if first != nil && prev.Val >= root.Val {
			second = root
		}

		prev = root

		traverse(root.Right)
	}

	prev = &TreeNode{ math.MinInt64, nil, nil }

	traverse(root)
	first.Val, second.Val = second.Val, first.Val
}
