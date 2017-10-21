package p054

func spiralOneLayer(m [][]int, rstart, cstart, rows, cols int) []int {
	switch {
	case rows == 1 && cols > 0:
		es := make([]int, cols)
		for ei, j := 0, cstart; j < cstart + cols; ei, j = ei+1, j+1 {
			es[ei] = m[rstart][j]
		}
		return es
	case rows > 0 && cols == 1:
		es := make([]int, rows)
		for ei, i := 0, rstart; i < rstart + rows; ei, i = ei+1, i+1 {
			es[ei] = m[i][cstart]
		}
		return es
	case rows == 1 && cols == 1:
		return []int { m[rstart][cstart] }
	}

	es := make([]int, 2*cols + (rows-2)*2)

	ei := 0
	for j := cstart; j < cstart + cols; ei, j = ei+1, j+1 {
		es[ei] = m[rstart][j]
	}

	ci := cstart + cols - 1
	for i := rstart+1; i < rstart + rows; ei, i = ei+1, i+1 {
		es[ei] = m[i][ci]
	}

	ri := rstart + rows - 1
	for j := cstart + cols - 2; j >= cstart; ei, j = ei+1, j-1 {
		es[ei] = m[ri][j]
	}

	for i := rstart + rows - 2; i > cstart; ei, i = ei+1, i-1 {
		es[ei] = m[i][cstart]
	}

	return es
}

func spiralOrder(matrix [][]int) []int {
	switch {
	case len(matrix) == 0:
		return []int {}
	case len(matrix) == 1:
		return matrix[0]
	case len(matrix[0]) == 1:
		es := make([]int, len(matrix))
		for ei, i := 0, 0; i < len(matrix); ei, i = ei+1, i+1 {
			es[ei] = matrix[i][0]
		}
		return es
	}

	es := []int {}
	rows, cols := len(matrix), len(matrix[0])
	for i,j := 0,0; i < (rows+1)/2 && j < (cols+1)/2; i, j = i+1, j+1 {
		es1 := spiralOneLayer(matrix, i, j, rows - 2*i, cols - 2*j)
		es = append(es, es1...)
	}
	return es
}
