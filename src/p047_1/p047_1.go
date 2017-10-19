package p047_1

import "sort"

func dfs(nums []int, used []bool, perm *([]int), res *([][]int)) {
	if len(*perm) == len(nums) {
		perm1 := make([]int, len(*perm))
		copy(perm1, *perm)

		*res = append(*res, perm1)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}

		if i > 0 && nums[i-1] == nums[i] && !used[i-1] {
			continue
		}

		used[i] = true
		*perm = append(*perm, nums[i])

		dfs(nums, used, perm, res)

		used[i] = false
		*perm = (*perm)[:len(*perm)-1]
	}
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)

	res := [][]int {}
	used := make([]bool, len(nums))

	dfs(nums, used, &([]int {}), &res)

	return res
}
