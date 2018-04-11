package p214_1

func reverseCopy(s string) string {
	n := len(s)
	if n == 0 {
		return s
	}
	s1 := make([]byte, n)
	for i := n - 1; i >= 0; i-- {
		s1[i] = s[n-1-i]
	}
	return string(s1)
}

func generateKMPTable(s string) []int {
	table := make([]int, len(s))

	idx := 0 // Pointer that points to matched char in prefix part

	// Skip index 0, so we will not match a string with itself
	for i := 1; i < len(s); i++ {
		if s[idx] == s[i] {
			// We can extend match in prefix and postfix
			table[i] = table[i-1] + 1
			idx++
		} else {
			// By assigning idx to table[i-1], we will shorten the match
			// string length, and jump to prefix part that we used to match
			// postfix ended at table[i-1]
			idx = table[i-1]

			for idx > 0 && s[idx] != s[i] {
				// We will try to shorten the length of match string
				// until we reach the beginning of match (idx 1)
				idx = table[idx-1]
			}

			// When we are here, either we found a match char or reached
			// the boundary without lack, so check again
			if s[idx] == s[i] {
				idx++
			}

			table[i] = idx
		}
	}

	return table
}

func shortestPalindrome(s string) string {
	s1 := s + "#" + reverseCopy(s)
	table := generateKMPTable(s1)

	ls := table[len(table)-1]
	return reverseCopy(s[ls:]) + s
}
