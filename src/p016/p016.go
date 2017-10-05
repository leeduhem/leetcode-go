package p016

import "sort"

func abs(x int) int {
	// Ignore possible overflow
	if (x < 0) {
		return -x
	}
	return x
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	nlen := len(nums)
	if nlen < 3 {
		panic("Invalid input")
	}

	cs := abs(nums[nlen-1]) + abs(nums[nlen-2]) + abs(nums[nlen-3])

	for i,vi := range nums {
		for j,vj := range nums[i+1:] {
			for _,vk := range nums[j+i+2:] {
				s := vi + vj + vk
				if abs(cs-target) >= abs(s-target) {
					cs = s
				}
			}
		}
	}

	return cs
}
