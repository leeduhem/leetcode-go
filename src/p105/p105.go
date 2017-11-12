package p105

func indexOf(nums []int, n int) int {
	// Find the index of the first occurrence
	for i, v := range nums {
		if v == n {
			return i
		}
	}

	return -1
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	val := preorder[0]

	if len(preorder) == 1 {
		return &TreeNode{ val, nil, nil }
	}

	vi := indexOf(inorder, val)

	left := buildTree(preorder[1 : vi+1], inorder[ : vi])
	right := buildTree(preorder[vi+1 : ], inorder[vi+1 : ])

	return &TreeNode{val, left, right}
}
