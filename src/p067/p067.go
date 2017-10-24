package p067

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func string2ReversedDigits(s string) []int {
	ds := make([]int, len(s))
	for i := len(s)-1; i >= 0; i-- {
		ds[i] = int( s[(len(s)-1)-i] - '0' )
	}
	return ds
}

func reversedDigits2String(ds []int) string {
	s := make([]rune, len(ds))
	for i := len(ds)-1; i >= 0; i-- {
		s[(len(ds)-1)-i] = rune ( ds[i] + '0' )
	}
	return string(s)
}

func addReversedDigits(a1, a2 []int, base int) []int {
	sum := make([]int, max(len(a1), len(a2)))

	add := func (a1, a2, c int) (int, int) {
		s := a1 + a2 + c
		if s >= base {
			s -= base
			c  = 1
		} else {
			c  = 0
		}
		return s,c
	}

	c, i := 0, 0
	for ; i < len(a1) && i < len(a2); i++ {
		sum[i], c = add(a1[i], a2[i], c)
	}

	for _, a := range [][]int {a1, a2} {
		for ; i < len(a); i++ {
			sum[i], c = add(a[i], 0, c)
		}
	}

	if c > 0 {
		return append(sum, c)
	}
	return sum
}

func addBinary(a string, b string) string {
	a1 := string2ReversedDigits(a)
	b1 := string2ReversedDigits(b)
	s  := addReversedDigits(a1, b1, 2)
	return reversedDigits2String(s)
}
