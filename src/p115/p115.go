package p115

func numDistinct(s string, t string) int {
	mem := make([][]int, len(t) + 1)
	for i, _ := range mem {
		mem[i] = make([]int, len(s) + 1)
	}

	// Filling the first row with 1s
	for j := 0; j <= len(s); j++ {
		mem[0][j] = 1
	}

	// The first column is 0 by default in every other rows except the first one,
	// that's what we need.

	for i := 0; i < len(t); i++ {
		for j := 0; j < len(s); j++ {
			if t[i] == s[j] {
				mem[i+1][j+1] = mem[i][j] + mem[i+1][j]
			} else {
				mem[i+1][j+1] = mem[i+1][j]
			}
		}
	}

	return mem[len(t)][len(s)]
}
