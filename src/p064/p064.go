package p064

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minPathSum(grid [][]int) int {
	type weightPair struct {
		Left, Up int
	}

	m, n := len(grid), len(grid[0])
	cm := make([][]weightPair, m)
	for i := 0; i < m; i++ {
		cm[i] = make([]weightPair, n)
	}

	cm[0][0] = weightPair{ grid[0][0], grid[0][0] }
	for j := 1; j < n; j++ {
		c, g := cm[0][j-1], grid[0][j]
		cm[0][j] = weightPair{ g + c.Left, g + c.Up }
	}
	for i := 1; i < m; i++ {
		c, g := cm[i-1][0], grid[i][0]
		cm[i][0] = weightPair{ g + c.Left, g + c.Up }
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			cu, cl := cm[i-1][j], cm[i][j-1]
			g := grid[i][j]
			cm[i][j] = weightPair{ min(cu.Left + g, cl.Left + g), min(cu.Up + g, cl.Up + g) }
		}
	}

	wp := cm[m-1][n-1]
	return min(wp.Left, wp.Up)
}
