package p179

import (
	"sort"
	"strconv"
	"strings"
)

func less(strs []string, i, j int) bool {
	s1 := strs[i] + strs[j]
	s2 := strs[j] + strs[i]
	return s1 < s2
}

func largestNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i, n := range nums {
		strs[i] = strconv.Itoa(n)
	}
	sort.Slice(strs, func(i, j int) bool { return !less(strs, i, j) })

	if strs[0][0] == '0' {
		return "0"
	}

	return strings.Join(strs, "")
}
