package p090

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	ss := [][]int { {} }
	for i := 0; i < len(nums); {
		cnt := 0
		for cnt + i < len(nums) && nums[cnt+i] == nums[i] {
			cnt++
		}

		prevN := len(ss)
		for j := 0; j < prevN; j++ {
			s := make([]int, len(ss[j]))
			copy(s, ss[j])
			for k := 0; k < cnt; k++ {
				s = append(s, nums[i])
				ss = append(ss, s)
			}
		}
		i += cnt
	}

	return ss
}
