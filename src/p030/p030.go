package p030

func findSubstring(s string, words []string) []int {
	k := len(words)
	inds := []int {}

	if k == 0 {
		return inds
	}

	wlen := len(words[0])
	slen := len(s)
	if k * wlen > slen {
		return inds
	}

	for i := 0; i <= slen - k * wlen; i++ {
		m := make(map[string]int, k)

		for j := i; j < i + k * wlen; j += wlen {
			key := s[j:j+wlen]
			m[key] += 1
		}

		for _,w := range words {
			if v,ok := m[w]; ok {
				if v == 1 {
					delete(m, w)
				} else {
					m[w] = v - 1
				}
			} else {
				break
			}
		}

		if len(m) == 0 {
			inds = append(inds, i)
		}
	}

	return inds
}
