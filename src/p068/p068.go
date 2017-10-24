package p068

import (
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	p := []string {}

	for wi := 0; wi < len(words); {
		line := []string { words[wi] }
		width := len(line[0])

		for width + len(line) - 1 < maxWidth && wi < len(words)-1 {
			wi++
			line = append(line, words[wi])
			width += len(words[wi])
		}

		if wi == len(words) - 1 && width + len(line) - 1 <= maxWidth {
			// Last line.
			l := strings.Join(line, " ")
			l += strings.Repeat(" ", maxWidth - len(l))
			p = append(p, l)
			break
		}

		if width + len(line) - 1 == maxWidth {
			wi++
			p = append(p, strings.Join(line, " "))
		} else {
			// Pop last word from current line.
			width -= len(words[wi])
			line = line[: len(line)-1]

			spaces, slots := (maxWidth - width), len(line) - 1
			if slots == 0 {
				l := line[0] + strings.Repeat(" ", maxWidth - len(line[0]))
				p = append(p, l)
				continue
			}

			sep, rem := spaces / slots, spaces % slots

			l := line[0]
			for i := 1; i < len(line); i++ {
				n := sep
				if rem > 0 {
					n++
					rem--
				}
				l += strings.Repeat(" ", n) + line[i]
			}

			p = append(p, l)
		}
	}

	return p
}
