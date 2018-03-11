package p171_1

func titleToNumber(s string) int {
	n := 0
	for _, c := range s {
		n = n * 26 + int(c) - int('A') + 1
	}
	return n
}
