package p202

func sumOfDigtSquares(n int) int {
	sum := 0
	for ; n > 0; n /= 10 {
		d := n % 10
		sum += d * d
	}
	return sum
}

func isHappy(n int) bool {
	mem := map[int]bool {n : true}

	for {
		n = sumOfDigtSquares(n)
		if n == 1 {
			return true
		}

		if _, ok := mem[n]; ok {
			return false
		} else {
			mem[n] = true
		}
	}

	return false
}
