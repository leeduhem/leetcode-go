package p204_1

func countPrimes(n int) int {
	notPrime := make([]bool, n)
	cnt := 0

	for i := 2; i < n; i++ {
		if ! notPrime[i] {
			cnt++
			for j := 2; i*j < n; j++ {
				notPrime[i*j] = true
			}
		}
	}

	return cnt
}
