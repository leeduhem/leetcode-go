package p069

func mySqrt(x int) int {
	for r := 0; ; r++ {
		if r*r > x {
			return r-1
		}
	}

	return 0
}
