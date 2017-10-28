package p079

func exist(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])

	used := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		used[i] = make([]bool, cols)
	}

	var search func(int, int, int) bool;
	search = func(rstart, cstart, wpos int) bool {
		if wpos == len(word) {
			return true
		}

		for _, cord := range []struct {
			row, col int
		} {
			{rstart-1, cstart},
			{rstart+1, cstart},
			{rstart, cstart-1},
			{rstart, cstart+1},
		} {
			r, c := cord.row, cord.col
			if r < 0 || r >= rows || c < 0 || c >= cols || used[r][c] {
				continue
			}
			if board[r][c] == word[wpos] {
				used[r][c] = true
				if found := search(r, c, wpos+1); found {
					return true
				}
				used[r][c] = false
			}
		}
		return false
	}


	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == word[0] {
				used[i][j] = true
				if found := search(i, j, 1); found {
					return true
				}
				used[i][j] = false
			}
		}
	}

	return false
}
