package p030_1

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

	ichan := make(chan int)

	for i := 0; i <= slen - k * wlen; i++ {
		go func(start int) {
			m := make(map[string]int, k)

			for j := start; j < start + k * wlen; j += wlen {
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
				ichan <- start
			} else {
				ichan <- -1
			}
		}(i)
	}

	for i := 0; i <= slen - k * wlen; i++ {
		select {
		case i := <- ichan:
			if i >= 0 {
				inds = append(inds, i)
			}
		}
	}

	return inds
}
