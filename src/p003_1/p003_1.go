// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
package p003_1

func intMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	m := make(map[rune]int, len(s))
	var j, max int

	for i, c := range s {
		if mc, ok := m[c]; ok {
			j = intMax(j, mc+1)
		}
		m[c] = i
		max = intMax(max, i-j+1)
	}

	return max
}
