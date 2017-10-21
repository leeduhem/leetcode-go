package p046_1

import "sort"

func dfs(nums []int, used []bool, perm []int, res *([][]int)) {
	if len(perm) == len(nums) {
		perm1 := make([]int, len(perm))
		copy(perm1, perm)

		*res = append(*res, perm1)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}

		used[i] = true
		dfs(nums, used, append(perm, nums[i]), res)
		used[i] = false
	}
}

func permute(nums []int) [][]int {
	sort.Ints(nums)

	res := [][]int {}
	used := make([]bool, len(nums))

	dfs(nums, used, []int {}, &res)
	return res
}
