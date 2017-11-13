package p106

func indexOf(nums []int, n int) int {
	// Find the index of the first occurrence
	for i, v := range nums {
		if v == n {
			return i
		}
	}

	return -1
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	val := postorder[len(postorder)-1]

	if len(postorder) == 1 {
		return &TreeNode{ val, nil, nil }
	}

	vi := indexOf(inorder, val)

	left := buildTree(inorder[ : vi], postorder[ : vi])
	right := buildTree(inorder[vi+1 : ], postorder[vi : len(postorder)-1])

	return &TreeNode{ val, left, right }
}
