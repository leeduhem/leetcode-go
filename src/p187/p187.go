package p187

func findRepeatedDnaSequences(s string) []string {
	w := map[int]int{}
	rs := make([]string, 0)

	m := make([]int, 26)
	m['C'-'A'] = 1
	m['G'-'A'] = 2
	m['T'-'A'] = 3

	for i := 0; i < len(s)-9; i++ {
		v := 0
		for j := i; j < i+10; j++ {
			v <<= 2
			v |= m[s[j]-'A']
		}
		w[v] += 1
		if c, _ := w[v]; c == 2 {
			rs = append(rs, s[i:i+10])
		}
	}

	return rs
}
