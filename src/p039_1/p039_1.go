package p039_1

import "sort"

func combinationSum1(candidates []int, target int, results *([][]int), combination []int, start int) {
	if target == 0 {
		// deep copy combination
		combination1 := make([]int, len(combination))
		copy(combination1, combination)

		*results = append(*results, combination1)
		return
	}
	for i := start; i < len(candidates) && target >= candidates[i]; i++ {
		combination = append(combination, candidates[i])
		combinationSum1(candidates, target-candidates[i], results, combination, i)
		combination = combination[:len(combination)-1]
	}
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	results := [][]int {}
	combinationSum1(candidates, target, &results, []int {}, 0)

	return results
}
