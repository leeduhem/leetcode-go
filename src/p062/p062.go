package p062

func uniquePaths(m int, n int) int {
	cm := make([][]int, m)
	for i,_ := range cm {
		cm[i] = make([]int, n)
	}

	for j := 0; j < n; j++ {
		cm[m-1][j] = 1
	}

	for i := 0; i < m; i++ {
		cm[i][n-1] = 1
	}

	for j := n-2; j >= 0; j-- {
		for i := m-2; i >= 0; i-- {
			cm[i][j] = cm[i][j+1] + cm[i+1][j]
		}
	}

	return cm[0][0]
}
