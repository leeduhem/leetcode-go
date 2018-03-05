package p164

import (
	"math"
)

func intMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func intMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maximumGap(nums []int) int {
	l := len(nums)
	if l < 2 {
		return 0
	}

	min, max := nums[0], nums[0]
	for _, n := range nums {
		min = intMin(min, n)
		max = intMax(max, n)
	}

	gap := int(math.Ceil(float64(max - min) / float64(l - 1)))
	buketsMin := make([]int, l-1)
	buketsMax := make([]int, l-1)
	for i := 0; i < l-1; i++ {
		buketsMin[i] = math.MaxInt32
		buketsMax[i] = math.MinInt32
	}

	for _, n := range nums {
		if n == min || n == max {
			continue
		}

		idx := (n - min) / gap
		buketsMin[idx] = intMin(n, buketsMin[idx])
		buketsMax[idx] = intMax(n, buketsMax[idx])
	}

	maxGap := math.MinInt32
	previous := min
	for i := 0; i < l-1; i++ {
		if buketsMin[i] == math.MaxInt32 && buketsMax[i] == math.MinInt32 {
			continue
		}

		maxGap = intMax(maxGap, buketsMin[i] - previous)
		previous = buketsMax[i]
	}
	maxGap = intMax(maxGap, max - previous)

	return maxGap
}
