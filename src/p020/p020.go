package p020

////////////////////////////////////////////////////////////////////////////////
// https://stackoverflow.com/questions/28541609/looking-for-reasonable-stack-implementation-in-golang

import (
	"sync"
	"errors"
)

type stack struct {
     lock sync.Mutex
     s []interface{}
}

func NewStack() *stack {
    return &stack {sync.Mutex{}, make([]interface{},0), }
}

func (s *stack) Push(v interface{}) {
    s.lock.Lock()
    defer s.lock.Unlock()

    s.s = append(s.s, v)
}

func (s *stack) Pop() (interface{}, error) {
    s.lock.Lock()
    defer s.lock.Unlock()


    l := len(s.s)
    if l == 0 {
        return 0, errors.New("Empty Stack")
    }

    res := s.s[l-1]
    s.s = s.s[:l-1]
    return res, nil
}

func (s *stack) IsEmpty() bool {
	return len(s.s) == 0
}

////////////////////////////////////////////////////////////////////////////////

func isOpenParenthese(c rune) bool {
	switch {
	case c == '(':
		fallthrough
	case c == '{':
		fallthrough
	case c == '[':
		return true
	default:
		return false
	}
}

func isCloseParenthese(c rune) bool {
	switch {
	case c == ')':
		fallthrough
	case c == '}':
		fallthrough
	case c == ']':
		return true
	default:
		return false
	}
}

func isMatch(p, p1 rune) bool {
	switch {
	case p == '(':
		return p1 == ')'
	case p == '{':
		return p1 == '}'
	case p == '[':
		return p1 == ']'
	default:
		return false
	}
}

func isValid(s string) bool {
	st := NewStack()

	for _,c := range s {
		if isOpenParenthese(c) {
			st.Push(c)
		} else if isCloseParenthese(c) {
			c1,err := st.Pop()
			if err != nil || !isMatch(c1.(rune), c) {
				return false
			}
		}
	}

	return st.IsEmpty()
}
