package p166

import (
	"strconv"
	"strings"
)

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func digit2str(d int) string {
	return strconv.Itoa(d)
}

func decimalParts(num, den int) (quot int, norep, rep []int) {
	quots := []int {}
	rems  := map[int]int {}

	for {
		q, r := num / den, num % den
		quots = append(quots, q)

		if r == 0 {
			return quots[0], quots[1:], []int {}
		}

		if p, ok := rems[r]; ok {
			return quots[0], quots[1:p], quots[p:]
		}

		rems[r] = len(quots)
		num = r * 10
	}

	// Should not reach here
	return
}

func fractionToDecimal(numerator int, denominator int) string {
	n1, n2, n3 := decimalParts(intAbs(numerator), intAbs(denominator))

	s2 := make([]string, len(n2))
	for i, n := range n2 {
		s2[i] = digit2str(n)
	}

	s3 := make([]string, len(n3))
	for i, n := range n3 {
		s3[i] = digit2str(n)
	}

	s := ""

	if (numerator < 0 && denominator > 0) || (numerator > 0 && denominator < 0) {
		s += "-"
	}

	s += strconv.Itoa(n1)

	if len(s2) > 0 || len(s3) > 0 {
		s += "."

		if len(s2) > 0 {
			s += strings.Join(s2, "")
		}
		if len(s3) > 0 {
			s += "(" + strings.Join(s3, "") + ")"
		}
	}
	return s
}
