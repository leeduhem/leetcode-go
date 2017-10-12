package p037

const dot = byte('.')

func findCandidate(board [][]byte) ([][]int) {
	cords := make([][]int, len(board[0]))

	for i,r := range board {
		cords1 := make([][]int, 0)
		for j,c := range r {
			if c == dot {
				cords1 = append(cords1, []int {i, j})
			}
		}

		if len(cords1) > 0 && len(cords1) < len(cords) {
			cords = cords1
		}
	}

	for j := 0; j < len(board[0]); j++ {
		cords1 := make([][]int, 0)
		for i,r := range board {
			if r[j] == dot {
				cords1 = append(cords1, []int {i, j})
			}
		}

		if len(cords1) > 0 && len(cords1) < len(cords) {
			cords = cords1
		}
	}

	for sr := 0; sr < 3; sr++ {
		for sc := 0; sc < 3; sc++ {
			cords1 := make([][]int, 0)
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					if board[3*sr+i][3*sc+j] == dot {
						cords1 = append(cords1, []int {3*sr+i, 3*sc+j})
					}
				}
			}

			if len(cords1) > 0 && len(cords1) < len(cords) {
				cords = cords1
			}
		}
	}

	for _,cord := range cords {
		if len(cord) != 0 {
			return cords
		}
	}

	return [][]int {}
}

// Is valid row, column, or sub-box?
func isValid(l []byte) bool {
	if len(l) != 9 {
		return false
	}

	m := make([]int, 10)
	for _,c := range l {
		if c != dot {
			k := int(c - '0')
			m[k] += 1
		}
	}

	for i := 1; i <= 9; i++ {
		if m[i] > 1 {
			return false
		}
	}

	return true
}

func isSolved(board [][]byte) bool {
	if ! isValidSudoku(board) {
		return false
	}

	for _,r := range board {
		for _,c := range r {
			if c == dot {
				return false
			}
		}
	}

	return true
}

func isValidSudoku(board [][]byte) bool {
	// check columns
	if len(board) != 9 {
		return false
	}
	for i := 0; i < 9; i++ {
		c := make([]byte, 9)
		for j,r := range board {
			c[j] = r[i]
		}

		if ok := isValid(c); !ok {
			return false
		}
	}

	// check rows
	for _,r := range board {
		if ok := isValid(r); !ok {
			return false
		}
	}

	// check sub-boxes
	for sr := 0; sr < 3; sr++ {
		for sc := 0; sc < 3; sc++ {
			s := make([]byte, 9)
			for r := 0; r < 3; r++ {
				for c := 0; c < 3; c++ {
					s[3*r+c] = board[3*sr+r][3*sc+c]
				}
			}
			if ok := isValid(s); !ok {
				return false
			}
		}
	}

	return true
}

func duplicateBoard(board [][]byte) [][]byte {
	b := make([][]byte, len(board))
	for i,r := range board {
		b[i] = make([]byte, len(r))
		copy(b[i], r)
	}
	return b
}

func fillMissedNumbers(board [][]byte, cords [][]int) ([][][]byte) {
	type numMap map[byte]int

	split := true
	numbers := make([]numMap, len(cords))

	for ci,cord := range cords {
		m := numMap {
			'1':1, '2':1, '3':1,
			'4':1, '5':1, '6':1,
			'7':1, '8':1, '9':1,
		}
		r, c := cord[0], cord[1]
		sr, sc := r/3, c/3

		// rows
		for _,k := range board[r] {
			delete(m, k)
		}
		if len(m) == 1 {
			goto Cont
		}

		// columns
		for _,row := range board {
			k := row[c]
			delete(m, k)
		}
		if len(m) == 1 {
			goto Cont
		}

		// sub-boxes
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				k := board[3*sr+i][3*sc+j]
				delete(m, k)
			}
		}
		if len(m) == 1 {
			goto Cont
		}

		numbers[ci] = m
		continue

	Cont:
		for k,_ := range m {
			board[r][c] = byte(k)
		}
		split = false
	}

	boards := [][][]byte { board }
	if split {
		// Generate all valid boards
		for i,m := range numbers {
			r, c := cords[i][0], cords[i][1]
			boards1 := [][][]byte {}
			for k,_ := range m {
				kb := [][][]byte {}
				for _,b := range boards {
					b1 := duplicateBoard(b)
					b1[r][c] = k
					if isValidSudoku(b1) {
						kb = append(kb, b1)
					}
				}
				boards1 = append(boards1, kb...)
			}

			boards = boards1
		}
	}

	return boards
}

func board2key(board [][]byte) string {
	str := ""
	for _,r := range board {
		str += string(r)
	}
	return str
}

var mem = map[string]bool {}

func solveSudoku(board [][]byte)  {
	for i := 0; i < 4; i++ {
		cords := findCandidate(board)
		if len(cords) == 0 {
			return
		}

		boards := fillMissedNumbers(board, cords)
		if len(boards) == 0 {
			return
		}

		for _,b := range boards {
			k := board2key(b)
			if _,ok := mem[k]; !ok {
				solveSudoku(b)
				if isSolved(b) {
					mem[k] = true

					for i,r := range b {
						copy(board[i], r)
					}

					return
				}
				mem[k] = false
			}
		}
	}
}
