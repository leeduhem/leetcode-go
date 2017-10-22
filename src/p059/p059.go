package p059

func fillOneLayer(m [][]int, rstart, cstart, rows, cols int, es *int) {
	e := *es
	switch {
	case rows == 1 && cols > 0:
		for j := cstart; j < cstart + cols; j++ {
			m[rstart][j] = e; e++
		}
		*es = e
		return
	case rows > 0 && cols == 1:
		for i := rstart; i < rstart + rows; i++ {
			m[i][cstart] = e; e++
		}
		*es = e
		return
	case rows == 1 && cols == 1:
		m[rstart][cstart] = e; e++
		*es = e
		return
	}

	for j := cstart; j < cstart + cols; j++ {
		m[rstart][j] = e; e++
	}

	ci := cstart + cols - 1
	for i := rstart+1; i < rstart + rows; i++ {
		m[i][ci] = e; e++
	}

	ri := rstart + rows - 1
	for j := cstart + cols - 2; j >= cstart; j-- {
		m[ri][j] = e; e++
	}

	for i := rstart + rows - 2; i > cstart; i-- {
		m[i][cstart] = e; e++
	}

	*es = e
	return
}

func generateMatrix(n int) [][]int {
	if n == 0 {
		return [][]int {}
	}
	if n == 1 {
		return [][]int { {1} }
	}

	// empty matrix
	matrix := make([][]int, n)
	for i,_ := range matrix {
		matrix[i] = make([]int, n)
	}

	h := (n+1)/2
	for es,i := 1,0; i < h; i++ {
		fillOneLayer(matrix, i, i, n-2*i, n-2*i, &es)
	}

	return matrix
}
