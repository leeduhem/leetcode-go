// https://leetcode.com/problems/regular-expression-matching/description/
package p010_1

import (
	"regexp"
)

func isMatch(s string, p string) bool {
	if matched, err := regexp.MatchString("^"+p+"$", s); err == nil && matched {
		return true
	}
	return false
}
