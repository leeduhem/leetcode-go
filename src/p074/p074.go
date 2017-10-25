package p074

import (
	"sort"
)

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rows, cols := len(matrix), len(matrix[0])

	col0 := make([]int, rows)
	for i,r := range matrix {
		col0[i] = r[0]
	}

	ri := sort.SearchInts(col0, target)
	if ri == 0 {
		if target < matrix[0][0] {
			return false
		}

		ci := sort.SearchInts(matrix[0], target)
		if ci == cols {
			return false
		}
		if matrix[0][ci] == target {
			return true
		}
		return false
	}

	if ri == rows {
		ci := sort.SearchInts(matrix[rows-1], target)
		if ci == cols {
			return false
		}
		if matrix[rows-1][ci] == target {
			return true
		}
		return false
	}

	for _, r := range []int {ri-1, ri} {
		ci := sort.SearchInts(matrix[r], target)
		if 0 <= ci && ci < cols && matrix[r][ci] == target {
			return true
		}
	}


	return false
}
