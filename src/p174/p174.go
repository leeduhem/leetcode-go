package p174

import (
	"math"
)

func intMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])

	hp := make([][]int, m+1)
	for i, _ := range hp {
		hp[i] = make([]int, n+1)
		for j, _ := range hp[i] {
			hp[i][j] = math.MaxInt32
		}
	}

	hp[m][n-1], hp[m-1][n] = 1, 1
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			need := intMin(hp[i+1][j], hp[i][j+1]) - dungeon[i][j]
			if need <= 0 {
				hp[i][j] = 1
			} else {
				hp[i][j] = need
			}
		}
	}

	return hp[0][0]
}
