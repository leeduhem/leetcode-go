package p031

import "sort"

// https://en.wikipedia.org/wiki/Permutation#Generation_in_lexicographic_order
func nextPermutation(nums []int)  {
	nlen := len(nums)

	k := -1
	for i := 0; i+1 < nlen; i++ {
		if nums[i] < nums[i+1] {
			k = i
		}
	}
	if k < 0 {
		sort.Ints(nums)
		return
	}

	l := k + 1
	for j := k+1; j < nlen; j++ {
		if nums[k] < nums[j] {
			l = j
		}
	}

	nums[k], nums[l] = nums[l], nums[k]

	// reverse nums[k+1, ..., nlen - 1]
	for i, j := k + 1, nlen - 1; i < j; i, j = i + 1, j - 1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
