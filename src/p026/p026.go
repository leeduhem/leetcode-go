package p026

func removeDuplicates(nums []int) int {
	nlen := len(nums)
	if nlen == 0 {
		return 0
	}

	lo := 0
	for hi := 1; hi < nlen; hi++ {
		for ; hi < nlen && nums[lo] == nums[hi]; hi++ {}

		if lo+1 < hi && hi < nlen {
			lo++
			nums[lo] = nums[hi]
		} else if lo+1 == hi {
			lo++
		}
	}

	return lo + 1
}
