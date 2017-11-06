package p097

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1) + len(s2) != len(s3) {
		return false
	}

	var t = make([][]bool, len(s1)+1)
	for i, _ := range t {
		t[i] = make([]bool, len(s2)+1)
	}

	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			if i == 0 && j == 0 {
				t[i][j] = true
			} else if i == 0 {
				t[i][j] = t[i][j-1] && s2[j-1] == s3[i+j-1]
			} else if j == 0 {
				t[i][j] = t[i-1][j] && s1[i-1] == s3[i-1+j]
			} else {
				t[i][j] = (t[i-1][j] && s1[i-1] == s3[i-1+j]) || (t[i][j-1] && s2[j-1] == s3[i+j-1])

			}
		}
	}

	return t[len(s1)][len(s2)]
}
