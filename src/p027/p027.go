package p027

func removeElement(nums []int, val int) int {
	nlen := len(nums)
	if nlen == 0 {
		return 0
	}

	lo := 0
	for hi := 0; hi < nlen; hi++ {
		for            ; lo < nlen && nums[lo] != val; lo++ {}
		for hi = lo + 1; hi < nlen && nums[hi] == val; hi++ {}

		if lo < hi && hi < nlen {
			nums[lo] = nums[hi]
			nums[hi] = val
			lo++
		}
	}

	return lo
}
