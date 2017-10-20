package p048

func rotate(matrix [][]int)  {
	matrix1 := make([][]int, len(matrix))
	for i,r := range matrix {
		matrix1[i] = make([]int, len(r))
	}

	rows := len(matrix)
	for j := rows-1; j >= 0; j-- {
		c := rows-1 - j
		for i,v := range matrix[j] {
			matrix1[i][c] = v
		}
	}

	for i,r := range matrix1 {
		copy(matrix[i], r)
	}
}
