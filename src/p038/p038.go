package p038

import "strconv"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	out := ""
	s := []rune(countAndSay(n-1))

	cur, cnt := s[0], 0
	for _,c := range s {
		if c == cur {
			cnt++
			continue
		}
		out += strconv.Itoa(cnt) + string([]rune{cur})
		cur, cnt = c, 1
	}
	out += strconv.Itoa(cnt) + string([]rune{cur})

	return out
}
