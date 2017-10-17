package p040

import "sort"

func combinationSum21(candidates []int, target int, results *([][]int), combination []int) {
	if target == 0 {
		// deep copy
		combination1 := make([]int, len(combination))
		copy(combination1, combination)

		*results = append(*results, combination1)
		return
	}

	for i := 0; i < len(candidates) && target >= candidates[i]; i++ {
		c := candidates[i]
		combinationSum21(candidates[i+1:], target-c, results, append(combination, c))

		// avoid duplicates
		for ; i+1 < len(candidates) && c == candidates[i+1]; i++ {}
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	results := [][]int {}
	combinationSum21(candidates, target, &results, []int {})

	return results
}
