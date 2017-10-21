package p051_2

import "strings"

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

	return true
}

func generateBoard(n int, qs []int) []string {
	// empty board
	b := make([]string, n)
	for i := 0; i < n; i++ {
		b[i] = strings.Repeat(".", n)
	}

	for i,q := range qs {
		b[i] = b[i][0:q] + "Q" + b[i][q+1:]
	}
	return b
}

func dfs(qs []int, used []bool, perm []int, bs *([][]string)) {
	n := len(qs)

	if len(perm) == n {
		*bs = append(*bs, generateBoard(n, perm))
		return
	}

	for i := 0; i < n; i++ {
		if used[i] {
			continue
		}

		perm1 := append(perm, qs[i])
		if ! isValidBoard(generateBoard(n, perm1)) {
			// Do not search further
			continue
		}

		used[i] = true
		dfs(qs, used, perm1, bs)
		used[i] = false
	}
}

func solveNQueens(n int) [][]string {
	qs := make([]int, n)
	for i := 0; i < n; i++ {
		qs[i] = i
	}

	bs := [][]string {}
	used := make([]bool, n)

	dfs(qs, used, []int {}, &bs)
	return bs
}
