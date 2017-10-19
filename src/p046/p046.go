package p046

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

func factoral(n int) int {
	f := 1
	for i := 1; i <= n; i++ {
		// ignore overflow
		f *= i
	}
	return f
}

func permute(nums []int) [][]int {
	sort.Ints(nums)

	nlen := len(nums)
	n := factoral(nlen)
	ps := make([][]int, n)
	for i := 0; i < n; i++ {
		p := make([]int, nlen)
		copy(p, nums)
		ps[i] = p

		nextPermutation(nums)
	}
	return ps
}
