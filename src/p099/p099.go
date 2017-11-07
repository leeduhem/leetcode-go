package p099

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


type cmpFunc func(int, int) bool

func recover(node, root *TreeNode, isViolation cmpFunc) {
	if root == nil {
		return
	}

	if isViolation(node.Val, root.Val) {
		node.Val, root.Val = root.Val, node.Val
		return
	}

	nv := node.Val

	recover(node, root.Left, isViolation)
	if nv != node.Val {
		return
	}

	recover(node, root.Right, isViolation)
}

func recoverSubtree(root *TreeNode)  {
	if root == nil {
		return
	}

	v := root.Val
	recover(root, root.Left, func (x, y int) bool { return x <= y })
	if root.Val != v {
		return
	}

	recover(root, root.Right, func (x, y int) bool { return x >= y })
	if root.Val != v {
		return
	}

	if root.Left != nil {
		v := root.Left.Val
		recoverSubtree(root.Left)
		if root.Left.Val != v {
			return
		}
	}

	if root.Right != nil {
		v := root.Right.Val
		recoverSubtree(root.Right)
		if root.Right.Val != v {
			return
		}
	}
}

func recoverTree(root *TreeNode) {
	if isValidBST(root) {
		return
	}

	for {
		recoverSubtree(root)
		if isValidBST(root) {
			return
		}
	}
}
