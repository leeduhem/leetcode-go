package p081_1

import (
	"sort"
)

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	// find pivot
	p := 0
	for ; p < len(nums); p++ {
		if p+1 < len(nums) && nums[p] > nums[p+1] {
			p++
			break
		}
	}

	if p <= 0 || p >= len(nums) {
		i := sort.SearchInts(nums, target)
		if (i == 0 && nums[i] == target) || (0 < i && i < len(nums) && nums[i] == target) {
			return true
		}
		return false
	}

	i := sort.SearchInts(nums[:p], target)
	if (i == 0 && nums[i] == target) || (0 < i && i < p && nums[i] == target) {
		return true
	}

	i = sort.SearchInts(nums[p:], target)
	if (i == 0 && nums[p] == target) || (0 < i && i < len(nums)-p && nums[i+p] == target) {
		return true
	}

	return false
}
