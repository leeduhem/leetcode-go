package p014

import (
	"math"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// Should be big enough
	minlen := math.MaxUint32

	for _,s := range strs {
		if len(s) < minlen {
			minlen = len(s)
		}
	}

	if minlen == 0 {
		return ""
	}

	maxlen := 0
OuterLoop:
	for l :=1 ; l <= minlen; l++ {
		p := strs[0][0:l]
		for _,s := range strs {
			if ! strings.HasPrefix(s, p) {
				break OuterLoop
			}
		}

		if maxlen < l {
			maxlen = l
		}
	}

	return strs[0][0:maxlen]
}
