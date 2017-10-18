package p044_1

func isMatch(s string, p string) bool {
	slen, plen := len(s), len(p)
	star, ss := -1, 0
	spos, ppos := 0, 0
	for spos < slen {
		// advancing both pointers when (both characters match) or ('?' is found in pattern)
		if ppos < plen && (p[ppos] == '?' || p[ppos] == s[spos]) {
			spos++; ppos++
			continue
		}

		// * found in pattern, track index of *, only advancing pattern pointer
		if ppos < plen && p[ppos] == '*' {
			star = ppos; ppos++
			ss = spos
			continue
		}

		// current character didn't match, last pattern pointer was '*', current pattern pointer is not '*'
		if star >= 0 {
			ppos = star + 1
			spos = ss + 1; ss++
			continue
		}

		// current pattern pointer is not star, last pattern pointer was not '*', character do not match
		return false
	}

	for ; ppos < plen && p[ppos] == '*'; ppos++ {}

	return ppos == plen && spos == slen
}
