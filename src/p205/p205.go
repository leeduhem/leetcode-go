package p205

func isIsomorphic(s string, t string) bool {
	var sm, tm [256]int

	for i := 0; i < len(s); i++ {
		sc, tc := int(s[i]), int(t[i])
		sm[sc], tm[tc] = i, i
	}

	for i := 0; i < len(s); i++ {
		sc, tc := int(s[i]), int(t[i])
		if sm[sc] != tm[tc] {
			return false
		}
	}
	return true
}
