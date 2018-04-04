package p209

import (
	"math"
)

func intMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minSubArrayLen(s int, nums []int) int {
	var (
		sum, from int
		win       int = math.MaxInt32
	)
	for i, num := range nums {
		sum += num
		for sum >= s {
			win = intMin(win, i-from+1)
			sum -= nums[from]
			from++
		}
	}

	if win == math.MaxInt32 {
		return 0
	}
	return win
}
