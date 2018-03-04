package p153

func findMin(nums []int) int {
	start := 0
	for end := len(nums)-1; start < end; {
		if nums[start] < nums[end] {
			return nums[start]
		}

		mid := (start+end) / 2

		if nums[mid] >= nums[start] {
			start = mid + 1
		} else {
			end = mid
		}
	}

	return nums[start]
}
