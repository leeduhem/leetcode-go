package p028

func strStr(haystack string, needle string) int {
	hlen := len(haystack)
	nlen := len(needle)

	if nlen == 0 {
		return 0
	}

	if hlen < nlen {
		return -1
	}

	for i := 0; i < hlen; i++ {
		if haystack[i] == needle[0] {
			j := 1
			for ; j < nlen; j++ {
				if i + j >= hlen {
					return -1
				}
				if i + j < hlen && haystack[i+j] != needle[j] {
					break
				}
			}

			if j == nlen {
				return i
			}
		}
	}

	return -1
}
