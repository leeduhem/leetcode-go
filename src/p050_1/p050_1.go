package p050_1

// See src/math/pow.go for the implementation of math.Pow(x, y float64) float64
func myPow(x float64, n int) float64 {
	if x == 1 || x == 0 {
		return x
	}

	if n == 0 {
		return 1
	}

	var n1 uint
	if n < 0 {
		// Ignore overflow
		n1 = uint(-n)
	} else {
		n1 = uint(n)
	}

	p := x
	c := uint(0)
	for i := n1; i > 1; i, c = i>>1, c+1 {
		p *= p
	}

	// Ignore overflow
	p *= myPow(x, int(n1 - (1<<c)))

	if n < 0 {
		return 1/p
	}

	return p
}
