package p034

// Looks like this log(N) implementation is fast enough for leetcode.
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int {-1, -1}
	}

	lo, hi := -1, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			lo = i
			break
		}
	}

	if lo < 0 {
		return []int {-1, -1}
	}

	for j := len(nums)-1; lo <= j; j-- {
		if nums[j] == target {
			hi = j
			break
		}
	}

	return []int {lo, hi}
}
