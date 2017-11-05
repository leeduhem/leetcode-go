package p094_1

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

func inorderTraversal(root *TreeNode) []int {
	nodes := []int {}
	seen := map[*TreeNode]bool {}

	subtrees := NewStack()
	subtrees.Push(root)

	for ! subtrees.IsEmpty() {
		v, ok := subtrees.Pop()
		if ok != nil {
			panic("Invalid input")
		}

		n := v.(*TreeNode)

		if n == nil {
			continue
		}

		if seen[n] {
			nodes = append(nodes, n.Val)
		} else {
			subtrees.Push(n.Right)
			subtrees.Push(n)
			subtrees.Push(n.Left)
		}

		seen[n] = true
	}

	return nodes
}
