package p051

import "strings"

type chessBoard []string

func isValidBoard(board chessBoard) bool {
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

var mem = map[string]bool {}

func addOneMoreQueen(b chessBoard) []chessBoard {
	bs := []chessBoard {}
	for i,r := range b {
		for j,c := range r {
			if c != 'Q' {
				// copy board
				b1 := make(chessBoard, len(b))
				copy(b1, b)

				b1[i] = b1[i][0:j] + "Q" + b1[i][j+1:]
				if isValidBoard(b1) {
					k := strings.Join(b1, ",")
					if _,ok := mem[k]; !ok {
						bs = append(bs, b1)
						mem[k] = true
					}
				}
			}
		}
	}
	return bs
}

func nQueens(n, q int) []chessBoard {
	if q == 0 {
		// empty chessboard
		b := make([]chessBoard, 1)
		b[0] = make(chessBoard, n)
		for i := 0; i < n; i++ {
			b[0][i] = strings.Repeat(".", n)
		}
		return b
	}

	bs := []chessBoard {}
	for _,b := range nQueens(n, q-1) {
		bs = append(bs, addOneMoreQueen(b)...)
	}
	return bs
}

func solveNQueens(n int) [][]string {
	// clear memory
	mem = map[string]bool {}

	bs := [][]string {}
	for _,b := range nQueens(n, n) {
		bs = append(bs, []string(b))
	}
	return bs
}
