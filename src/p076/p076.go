package p076

import (
	"math"
)

func minWindow(s string, t string) string {
	m := map[byte]int {}
	for i := 0; i < len(t); i++ {
		m[t[i]]++
	}

	dist, head := math.MaxInt32, 0
	for cnt, begin, end := len(t), 0, 0; end < len(s); {
		c := s[end]
		if m[c] > 0 {
			cnt-- // in t
		}
		end++; m[c]--

		for cnt == 0 { // valid
			if end - begin < dist {
				head = begin
				dist = end - head
			}
			c := s[begin]
			if m[c] == 0 {
				cnt++ // make it invalid
			}
			begin++; m[c]++
		}
	}

	if dist == math.MaxInt32 {
		return ""
	}
	return s[head : head+dist]
}
