package p063

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	cm := make([][]int, m)
	for i := 0; i < m; i++ {
		cm[i] = make([]int, n)
	}

	for o, j := 0, n-1; j >= 0; j-- {
		if obstacleGrid[m-1][j] != 0 {
			o = 1
		}
		cm[m-1][j] = 1 - o
	}

	for o, i := 0, m-1; i >= 0; i-- {
		if obstacleGrid[i][n-1] != 0 {
			o = 1
		}
		cm[i][n-1] = 1 - o
	}

	for j := n-2; j >= 0; j-- {
		for i := m-2; i >= 0; i-- {
			if obstacleGrid[i][j] != 0 {
				cm[i][j] = 0
			} else {
				cm[i][j] = cm[i+1][j] + cm[i][j+1]
			}
		}
	}

	return cm[0][0]
}
