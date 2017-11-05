package p095

func generate(vs []int) []*TreeNode {
	trees := []*TreeNode {}

	for i, v := range vs {
		lefts  := generate(vs[:i])
		rights := generate(vs[i+1:])

		if len(lefts) == 0 {
			lefts = []*TreeNode { nil }
		}
		if len(rights) == 0 {
			rights = []*TreeNode { nil }
		}

		for _, left := range lefts {
			for _, right := range rights {
				trees = append(trees, &TreeNode{ v, left, right })
			}
		}
	}

	return trees
}

func generateTrees(n int) []*TreeNode {
	vs := make([]int, n)
	for i := 0; i < n; i++ {
		vs[i] = i+1
	}

	return generate(vs)
}
