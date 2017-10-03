package p007

import (
	"strconv"
	"unicode"
	"math"
)

func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func reverse(x int) int {
	xs := strconv.Itoa(x)

	var xs1 string
	if (unicode.IsDigit(rune(xs[0]))) {
		xs1 = reverseString(xs)
	} else {
		xs1 = xs[0:1] + reverseString(xs[1:])
	}

	x1,err := strconv.Atoi(xs1)
	if (err == nil) {
		if (math.MaxInt32 < x1 || x1 < math.MinInt32) {
			return 0
		}
		return x1
	} else {
		return 0
	}
}

