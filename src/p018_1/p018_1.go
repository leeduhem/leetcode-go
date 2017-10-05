package p018_1

import "sort"

func fourSum(nums []int, target int) [][]int {
	res := [][]int {}

	nlen := len(nums)
	if (nlen < 4) {
		return res
	}

	sort.Ints(nums)
	max := nums[nlen-1]
	if 4 * nums[0] > target || 4 * max < target {
		return res
	}

	for i,n := range nums {
		// avoid duplicates
		if i > 0 && n == nums[i-1] {
			continue
		}
		// n is too small
		if n + 3 * max < target {
			continue
		}
		// n is too large
		if 4 * n > target {
			break
		}
		// n is the boundary
		if 4 * n == target {
			if i + 3 < nlen && nums[i+3] == n {
				res = append(res, []int{n, n, n, n})
				}
			break
		}

		res = append(res, threeSum(nums[i+1:], target - n, n)...)
	}

	return res
}

func threeSum(nums []int, target, n1 int) [][]int {
	res := [][]int {}

	nlen := len(nums)
	if nlen < 3 {
		return res
	}

	max := nums[nlen-1]
	if 3 * nums[0] > target || 3 * max < target {
		return res
	}

	for i,n := range nums {
		if i > 0 && n == nums[i-1] {
			continue
		}
		if n + 2 * max < target {
			continue
		}
		if 3 * n > target {
			break
		}

		if 3 * n == target {
			if i + 2 < nlen && nums[i+2] == n {
				res = append(res, []int {n1, n, n, n})
			}
			break
		}

		res = append(res, twoSum(nums[i+1:], target - n, n1, n)...)
	}

	return res
}

func twoSum(nums []int, target, n1, n2 int) [][]int {
	res := [][]int {}

	nlen := len(nums)
	if nlen < 2 {
		return res
	}

	if 2 * nums[0] > target || 2 * nums[nlen-1] < target {
		return res
	}

	for lo,hi := 0,nlen-1; lo < hi; {
		sum := nums[lo] + nums[hi]
		if sum == target {
			res = append(res, []int{n1, n2, nums[lo], nums[hi]})

			for ; lo < hi && nums[lo] == nums[lo+1]; lo++ {}
			for ; lo < hi && nums[hi] == nums[hi-1]; hi-- {}
			lo++; hi--
		} else if sum < target {
			lo++
		} else if sum > target {
			hi--
		}
	}

	return res
}
