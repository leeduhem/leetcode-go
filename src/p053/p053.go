package p053

func maxSubArray(nums []int) int {
	// Should be small enough.
	max := -1 << 31

	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum > max {
				max = sum
			}
		}
	}

	return max
}
