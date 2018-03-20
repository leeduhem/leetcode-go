package p199

func rightView(node *TreeNode, result *[]int, deep int) {
	if node == nil {
		return
	}

	if deep == len(*result) {
		*result = append(*result, node.Val)
	}

	rightView(node.Right, result, deep+1)
	rightView(node.Left, result, deep+1)
}

func rightSideView(root *TreeNode) []int {
	result := new([]int)
	rightView(root, result, 0)
	return *result
}
