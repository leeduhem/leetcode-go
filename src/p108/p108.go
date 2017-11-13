package p108

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &TreeNode{ nums[0], nil, nil }
	}

	mid := len(nums) / 2

	val := nums[mid]
	left := sortedArrayToBST(nums[ : mid])
	right := sortedArrayToBST(nums[mid+1 : ])

	return &TreeNode{ val, left, right }
}
