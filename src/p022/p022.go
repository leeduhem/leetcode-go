package p022

import "strings"

func generateParenthesis(n int) []string {
	type part struct {
		oc, so int
		p string
	}

	ps := []part { part{0, 0, ""} }

	for {
		cont := false
		ps1 := make([]part, 0)

		for _,p := range ps {
			if p.oc == n {
				p1 := p
				if p.so > 0 {
					p1 = part{ p.oc, 0, p.p + strings.Repeat(")", p.so)}
				}

				ps1 = append(ps1, p1)
				continue
			}

			if p.so == 0 {
				ps1 = append(ps1, part{ p.oc + 1, p.so + 1, p.p + "("})
				cont = true
				continue
			}

			ps1 = append(ps1, part{ p.oc + 1, p.so + 1, p.p + "(" }, part{ p.oc, p.so - 1, p.p + ")" })
			cont = true
		}

		ps = ps1

		if !cont {
			break
		}
	}

	pstrs := make([]string, len(ps))
	for i,p := range ps {
		pstrs[i] = p.p
	}

	return pstrs
}
