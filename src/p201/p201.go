package p201

func rangeBitwiseAnd(m int, n int) int {
	factor := 1
	for m != n {
		m >>= 1
		n >>= 1
		factor <<= 1
	}
	return m * factor
}
