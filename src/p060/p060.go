package p060

import (
	"sort"
	"strconv"
	"strings"
)

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

func ints2str(ns []int) string {
	ss := make([]string, len(ns))
	for i,n := range ns {
		ss[i] = strconv.Itoa(n)
	}
	return strings.Join(ss, "")
}

func getPermutation(n, k int) string {
	ns := make([]int, n)
	for i,_ := range ns {
		ns[i] = i+1
	}

	for i := 0; i < k-1; i++ {
		nextPermutation(ns)
	}

	return ints2str(ns)
}
