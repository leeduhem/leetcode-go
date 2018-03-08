package p167

import (
	"sort"
)

func twoIndexes(numbers []int, target int) (int, int) {
	for idx, num := range numbers {
		numbers1 := numbers[idx + 1 :]
		target1  := target - num

		idx1 := sort.SearchInts(numbers1, target1)
		if idx1 < len(numbers1) && numbers1[idx1] == target1 {
			return idx, idx + 1 + idx1
		}
	}
	return -1, -1
}

func twoSum(numbers []int, target int) []int {
	idx1, idx2 := twoIndexes(numbers, target)
	if idx1 < 0 {
		return []int{-1, -1}
	}
	return []int{idx1 + 1, idx2 + 1}
}
