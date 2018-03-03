package p150

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, token := range tokens {
		if n, err := strconv.Atoi(token); err == nil {
			stack = append(stack, n)
		} else {
			l := len(stack)
			n1, n2 := stack[l-1], stack[l-2]

			switch token {
			case "+":
				stack = append(stack[:l-2], n1+n2)
			case "-":
				stack = append(stack[:l-2], n2 - n1)
			case "*":
				stack = append(stack[:l-2], n1 * n2)
			case "/":
				stack = append(stack[:l-2], n2 / n1)
			}
		}
	}

	return stack[0]
}
