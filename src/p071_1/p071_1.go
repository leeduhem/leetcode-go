package p071_1

import (
	"strings"
)

func simplifyPath(p string) string {
	comps := strings.Split(p, "/")

	// clean components
	comps1 := []string {}
	for _,c := range comps {
		if c == "." || c == "" {
			continue
		}
		if c == ".." {
			if len(comps1) > 0 {
				comps1 = comps1[: len(comps1)-1]
			}
			continue
		}
		comps1 = append(comps1, c)
	}

	return "/" + strings.Join(comps1, "/")
}
