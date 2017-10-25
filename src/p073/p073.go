package p073

func setZeroes(matrix [][]int)  {
	if len(matrix) == 0 {
		return
	}

	m, n := len(matrix), len(matrix[0])

	erows := []int {}
	for i, r := range matrix {
		for _, c := range r {
			if c == 0 {
				erows = append(erows, i)
				break
			}
		}
	}

	ecols := []int {}
	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			if matrix[i][j] == 0 {
				ecols = append(ecols, j)
				break
			}
		}
	}

	for _, i := range erows {
		for j := 0; j < n; j++ {
			matrix[i][j] = 0
		}
	}

	for _, j := range ecols {
		for i := 0; i < m; i++ {
			matrix[i][j] = 0
		}
	}
}
