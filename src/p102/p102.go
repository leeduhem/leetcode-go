package p102

func levelOrder(root *TreeNode) [][]int {
	valuesList := [][]int {}

	subtrees := []*TreeNode { root }
	for len(subtrees) > 0 {
		values := []int {}
		subtrees1 := []*TreeNode {}
		for _, tree := range subtrees {
			if tree == nil {
				continue
			}

			values = append(values, tree.Val)
			subtrees1 = append(subtrees1, tree.Left)
			subtrees1 = append(subtrees1, tree.Right)
		}

		if len(values) > 0 {
			valuesList = append(valuesList, values)
		}
		subtrees = subtrees1
	}

	return valuesList
}
