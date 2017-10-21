package p052_1

func backtracking(row int, cols, d1, d2 []bool, n, cnt int) int {
	if row == n {
		cnt++
		return cnt
	}

	for col := 0; col < n; col++ {
		id1 := col - row + n
		id2 := col + row
		if cols[col] || d1[id1] || d2[id2] {
			continue
		}

		cols[col], d1[id1], d2[id2] = true, true, true
		cnt = backtracking(row+1, cols, d1, d2, n, cnt)
		cols[col], d1[id1], d2[id2] = false, false, false
	}

	return cnt
}

func totalNQueens(n int) int {
	cols := make([]bool, n)	  // columns
	d1   := make([]bool, 2*n) // diagonals \
	d2   := make([]bool, 2*n) // diagonals /
	return backtracking(0, cols, d1, d2, n, 0)
}
