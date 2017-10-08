package p029_1

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
	if dividend <= math.MinInt32 && divisor == -1 {
		return math.MaxInt32 // overflow
	}

	sign := 1
	if (dividend < 0) != (divisor < 0) {
		sign = -1
	}

	dividend1 := abs(int64(dividend))
	divisor1  := abs(int64(divisor))

	quote := 0
	for dividend1 >= divisor1 {
		t, m := divisor1, 1
		for dividend1 >= (t << 1) {
			t <<= 1
			m <<= 1
		}

		dividend1 -= t
		quote += m
	}

	// overflow
	if sign < 0 && quote > math.MaxInt32 {
		return math.MinInt32
	}

	return sign * quote
}
