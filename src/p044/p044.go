package p044

import "path"

// cheating...
func isMatch(s string, p string) bool {
	matched, err := path.Match(p, s)
	if err != nil {
		return false
	}
	return matched
}
