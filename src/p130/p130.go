package p130

func solve(board [][]byte) {
	row := len(board)
	if row == 0 {
		return
	}
	col := len(board[0])

	if col <= 2 || row <= 2 {
		return
	}

	var alter func(i, j int)
	alter = func(i, j int) {
		if board[i][j] != 'O' {
			return
		}

		board[i][j] = '1'
		if i > 1 {
			alter(i-1, j)
		}
		if j > 1 {
			alter(i, j-1)
		}
		if i+1 < row {
			alter(i+1, j)
		}
		if j+1 < col {
			alter(i, j+1)
		}
	}

	for i := 0; i < row; i++ {
		alter(i, 0)
		alter(i, col-1)
	}
	for j := 1; j+1 < col; j++ {
		alter(0, j)
		alter(row-1, j)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '1' {
				board[i][j] = 'O'
			}
		}
	}
}
