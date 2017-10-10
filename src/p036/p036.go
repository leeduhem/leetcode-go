package p036

const dot = byte('.')

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
