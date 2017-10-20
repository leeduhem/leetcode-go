package p048_1

func rotate(matrix [][]int)  {
	// reverse up to down
	for i := len(matrix)/2-1; i >= 0; i-- {
		opp := len(matrix)-1-i
		matrix[i], matrix[opp] = matrix[opp], matrix[i]
	}

	// swap the symmetry
	for i := 0; i < len(matrix); i++ {
		for j := i+1; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
