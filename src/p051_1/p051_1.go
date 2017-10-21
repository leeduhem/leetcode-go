package p051_1

import (
	"sort"
	"strings"
)

// https://en.wikipedia.org/wiki/Permutation#Generation_in_lexicographic_order
func nextPermutation(nums []int)  {
	nlen := len(nums)

	k := -1
	for i := 0; i+1 < nlen; i++ {
		if nums[i] < nums[i+1] {
			k = i
		}
	}
	if k < 0 {
		sort.Ints(nums)
		return
	}

	l := k + 1
	for j := k+1; j < nlen; j++ {
		if nums[k] < nums[j] {
			l = j
		}
	}

	nums[k], nums[l] = nums[l], nums[k]

	// reverse nums[k+1, ..., nlen - 1]
	for i, j := k + 1, nlen - 1; i < j; i, j = i + 1, j - 1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func factoral(n int) int {
	f := 1
	for i := 1; i <= n; i++ {
		// ignore overflow
		f *= i
	}
	return f
}

func isValidBoard(board []string) bool {
	rows, cols := len(board), len(board[0])

	// check diagonals
	for j := 0; j < cols; j++ {
		q := 0
		for i1, j1 := 0, j; i1 < rows && j1 < cols; i1, j1 = i1+1, j1+1 {
			if board[i1][j1] == 'Q' {
				q++
				if q > 1 { return false }
			}
		}
	}
	for i := 1; i < rows; i++ {
		q := 0
		for i1, j1 := i, 0; i1 < rows && j1 < cols; i1, j1 = i1+1, j1+1 {
			if board[i1][j1] == 'Q' {
				q++
				if q > 1 { return false }
			}
		}
	}

	for j := 0; j < cols; j++ {
		q := 0
		for i1, j1 := rows-1, j; i1 >= 0 && j1 < cols; i1, j1 = i1-1, j1+1 {
			if board[i1][j1] == 'Q' {
				q++
				if q > 1 { return false }
			}
		}
	}
	for i := rows-2; i >= 0; i-- {
		q := 0
		for i1, j1 := i, 0; i1 >= 0 && j1 < cols; i1, j1 = i1-1, j1+1 {
			if board[i1][j1] == 'Q' {
				q++
				if q > 1 { return false }
			}
		}
	}

	// check rows
	for _,r := range board {
		q := 0
		for _,c := range r {
			if c == 'Q' {
				q++
				if q > 1 { return false }
			}
		}
	}

	// check columns
	for j := 0; j < cols; j++ {
		q := 0
		for i := 0; i < rows; i++ {
			if board[i][j] == 'Q' {
				q++
				if q > 1 { return false }
			}
		}
	}

	return true
}

func generateBoard(qs []int) []string {
	qlen := len(qs)
	t := strings.Repeat(".", qlen)

	b := make([]string, qlen)
	for i,q := range qs {
		b[i] = t[0:q] + "Q" + t[q+1:]
	}
	return b
}

func solveNQueens(n int) [][]string {
	qs := make([]int, n)
	for i := 0; i < n; i++ {
		qs[i] = i
	}

	bs := [][]string {}
	for i := 0; i < factoral(n); i++ {
		b := generateBoard(qs)
		if isValidBoard(b) {
			bs = append(bs, b)
		}
		nextPermutation(qs)
	}

	return bs
}
