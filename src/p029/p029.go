package p029

import "math"

func abs(x int64) int64 {
	// overflow
	if x == math.MinInt64 {
		return math.MaxInt64
	}

	if x < 0 {
		return -x
	}

	return x
}

func divide(dividend int, divisor int) int {
	if divisor == 0 {
		panic("Dividing by zero")
	}

	if divisor == 1 {
		return dividend
	}

	// Seems that leetcode is using 32-bit machine
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32 // overflow
	}

	isNegative := false
	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		isNegative = true
	}

	dividend1 := abs(int64(dividend))
	divisor1  := abs(int64(divisor))

	quote := 0
	for dividend1 >= divisor1 {
		dividend1 -= divisor1
		quote++
	}

	if isNegative {
		// overflow
		if quote > math.MaxInt32 {
			return math.MinInt32
		}
		return -quote
	}
	return quote
}
