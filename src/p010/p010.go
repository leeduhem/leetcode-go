// https://leetcode.com/problems/regular-expression-matching/description/
package p010

// NB. Real regexp engine is far more complex than this,
// and this is NOT how they are implemented.

func isMatch(s string, p string) bool {
	strlen, patlen := len(s), len(p)
	strpos, patpos := 0, 0

	for strpos < strlen && patpos < patlen {
		zeroMore := false
		if patpos < patlen-1 && p[patpos+1] == '*' {
			zeroMore = true
		}

		if p[patpos] == '.' || p[patpos] == s[strpos] {
			strpos++
			if !zeroMore {
				patpos++
			}
		} else if zeroMore {
			patpos += 2
		} else {
			break
		}
	}

	if strpos == 0 {
		if strlen != 0 {
			return false
		}

		if patlen == 0 {
			return true
		}

		// Ignore further ZeroOrMore patterns
		for ; patpos < patlen-1 && p[patpos+1] == '*'; patpos += 2 {
		}
		if patpos == patlen {
			return true
		}

		return false
	}

	if strpos == strlen {
		// Ignore further ZeroOrMore patterns
		for ; patpos < patlen-1 && p[patpos+1] == '*'; patpos += 2 {
		}
		if patpos == patlen {
			return true
		}
	}

	// backtrack
	if 0 < patpos && patpos < patlen {
		for strpos1 := strlen - 1; strpos1 >= 0; strpos1-- {
			if isMatch(s[0:strpos1], p[0:patpos]) && isMatch(s[strpos1:], p[patpos:]) {
				return true
			}
		}

		for patpos1 := patpos - 2; patpos1 >= 0; {
			if isMatch(s[0:strpos], p[0:patpos1]) && isMatch(s[strpos:], p[patpos1:]) {
				return true
			}

			if patpos1 > 1 && p[patpos1-1] == '*' {
				patpos1 -= 2
			} else {
				patpos1 -= 1
			}
		}
	}

	return false
}
