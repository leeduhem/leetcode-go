package p065

import (
	"unicode"
)

const (
	space = iota
	sign
	digit
	dot
	exponent

	invalidState
)

func isValidState(state int, validStates []int) bool {
	for _,st := range validStates {
		if state == st {
			return true
		}
	}
	return false
}

// This is NOT how finite state machine (FSM) is implemented in general.
func isNumber(s string) bool {
	validStates := []int {space, sign, digit, dot}
	newValidStates := []int {}

	atBeginning := true
	haveDigit   := false
	needMoreDigit := false
	haveDot     := false
	haveExponent := false
	state       := invalidState

	for _,c := range s {
		switch {
		case unicode.IsSpace(c):
			state = space
			if atBeginning {
				newValidStates = validStates
			} else {
				newValidStates = []int {space}
			}
		case c == '+' || c == '-':
			state = sign
			if haveExponent {
				newValidStates = []int {digit}
			} else if ! haveDigit {
				newValidStates = []int {digit, dot}
			} else {
				return false
			}
			atBeginning = false
		case unicode.IsDigit(c):
			state = digit
			if haveDot {
				newValidStates = []int {digit, exponent, sign, space}
			} else {
				newValidStates = []int {digit, dot, exponent, sign, space}
			}
			if haveExponent {
				newValidStates = []int {digit, space}
			}
			haveDigit = true
			atBeginning = false
			needMoreDigit = false
		case c == '.':
			state = dot
			if haveExponent {
				return false
			}
			if haveDigit {
				newValidStates = []int {digit, exponent, sign, space}
			} else {
				newValidStates = []int {digit}
			}
			haveDot = true
			atBeginning = false
		case c == 'e' || c == 'E':
			state = exponent
			if ! haveDigit || haveExponent {
				return false
			}
			newValidStates = []int {sign, digit}
			atBeginning = false
			haveExponent = true
			needMoreDigit = true
		default:
			return false
		}

		if ! isValidState(state, validStates) {
			return false
		}
		validStates = newValidStates
	}

	if haveDigit && !needMoreDigit {
		return true
	}
	return false
}
