package p087

func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}

	l1 := len(s1)
	if l1 != len(s2) {
		return false
	}

	m := map[byte]int {}
	for i := 0; i < l1; i++ {
		m[s1[i]] += 1
		m[s2[i]] -= 1
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	for i := 1; i < l1; i++ {
		if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) {
			return true
		}
		if isScramble(s1[:i], s2[l1-i:]) && isScramble(s1[i:], s2[:l1-i]) {
			return true
		}
	}

	return false
}
