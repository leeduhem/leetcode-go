package p204

import (
	"math"
)

func isPrime(n int) bool {
	if n % 2 == 0 {
		return false
	}

	m := int(math.Sqrt(float64(n)))
	for f := 3; f <= m; f += 2 {
		if n % f == 0 {
			return false
		}
	}
	return true
}

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}

	cnt := 1
	for i := 3; i < n; i += 2 {
		if isPrime(i) {
			cnt++
		}
	}
	return cnt
}
