// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
package p003

func lengthOfLongestSubstring(s string) int {
	max := 0
	if len(s) > 0 {
		max = 1
	}

	for i, vi := range s {
		m := make(map[rune]int, 26)
		m[vi] = i

		for j, vj := range s[i+1:] {
			if _, ok := m[vj]; ok {
				break
			} else {
				m[vj] = j + i + 1
			}

			// (j + i + 1) - i + 1 = j + 2
			if j+2 > max {
				max = j + 2
			}
		}
	}

	return max
}
