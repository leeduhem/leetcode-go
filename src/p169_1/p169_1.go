package p169_1

func majorityElement(nums []int) int {
	major, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			count++
			major = nums[i]
		} else if major == nums[i] {
			count++
		} else {
			count--
		}
	}

	return major
}
