package p123

import (
	"math"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxProfit(prices []int) int {
	hold1, hold2 := math.MinInt32, math.MinInt32
	release1, release2 := 0, 0
	for _, p := range prices {
		release2 = max(release2, hold2 + p)
		hold2    = max(hold2, release1 - p)
		release1 = max(release1, hold1 + p)
		hold1    = max(hold1, -p)
	}
	return release2
}
