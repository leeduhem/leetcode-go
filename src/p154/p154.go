package p154

func findMin(nums []int) int {
	start := 0
	for end := len(nums)-1; start < end; {
		mid := (start+end) / 2

		if nums[mid] > nums[end] {
			start = mid + 1
		} else if nums[mid] < nums[end] {
			end = mid
		} else {
			end--
		}
	}

	return nums[start]
}
