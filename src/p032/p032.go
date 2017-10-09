package p032

func longestValidParentheses(s string) int {
	top := -1
	st  := make([]int, len(s))

	for i,c := range s {
		if c == '(' {
			top++
			st[top] = i
		} else if top >= 0 && s[st[top]] == '(' {
			top--
		} else {
			top++
			st[top] = i
		}
	}

	if top < 0 {
		return len(s)
	}

	longest := 0
	a := len(s)
	for b := 0; top >= 0; {
		b = st[top]; top--
		if a - b - 1 > longest {
			longest = a - b -1
		}
		a = b
	}
	if a > longest {
		longest = a
	}

	return longest
}
