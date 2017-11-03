package p091

import (
	"strconv"
)

func numDecodings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	mem := make([]int, n+1)
	mem[n] = 1
	if s[n-1] == '0' {
		mem[n-1] = 0
	} else {
		mem[n-1] = 1
	}

	for i := n-2; i >= 0; i-- {
		if s[i] == '0' {
			continue
		} else {
			if c, ok := strconv.Atoi(s[i:i+2]); ok == nil && c <= 26 {
				mem[i] = mem[i+1] + mem[i+2]
			} else {
				mem[i] = mem[i+1]
			}
		}
	}

	return mem[0]
}
