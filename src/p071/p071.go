package p071

import (
	"path"
)

// Cheating...
func simplifyPath(p string) string {
	return path.Clean(p)
}
