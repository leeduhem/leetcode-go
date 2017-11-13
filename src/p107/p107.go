package p107

func reverse(a [][]int) {
	for i := len(a)/2-1; i >= 0; i-- {
		opp := len(a)-1-i
		a[i], a[opp] = a[opp], a[i]
	}
}

func levelOrderBottom(root *TreeNode) [][]int {
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

	reverse(valuesList)

	return valuesList
}
