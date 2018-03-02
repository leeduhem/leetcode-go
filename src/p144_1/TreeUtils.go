package p144_1

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isTreeEqual(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}

	if t1 != nil && t2 != nil {
		return t1.Val == t2.Val &&
			isTreeEqual(t1.Left, t2.Left) &&
			isTreeEqual(t1.Right, t2.Right)
	}

	return false
}

func deepCopyTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	return &TreeNode{ root.Val, deepCopyTree(root.Left), deepCopyTree(root.Right) }
}
