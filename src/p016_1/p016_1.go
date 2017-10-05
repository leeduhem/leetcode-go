package p016_1

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

	csum := abs(nums[nlen-1]) + abs(nums[nlen-2]) + abs(nums[nlen-3])

	for i := 0; i < nlen - 2; i++ {
		lo,hi := i+1, nlen-1

		for ; lo < hi; {
			sum := nums[i] + nums[lo] + nums[hi]
			if abs(csum-target) >= abs(sum-target) {
				csum = sum
				if csum == target {
					return csum
				}
			}
			if sum > target {
				hi--
			} else {
				lo++
			}
		}
	}

	return csum
}
