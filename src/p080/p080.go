package p080

func removeDuplicates(nums []int) int {
	nlen := len(nums)
	if nlen <= 2 {
		return nlen
	}

	l := 0
	for r := 1; r < len(nums); {
		if nums[l] == nums[r] && l > 0 && nums[l-1] == nums[l] {
			r++
			continue
		}

		l++
		if l < r { nums[l] = nums[r] }
		r++
	}

	return l+1
}
