package p101_1

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

func (s *stack) Size() int {
	return len(s.s)
}

////////////////////////////////////////////////////////////////////////////////


func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil || root.Right == nil {
		return root.Left == root.Right
	}

	subtrees := NewStack()
	subtrees.Push(root.Left)
	subtrees.Push(root.Right)

	for ! subtrees.IsEmpty() {
		if subtrees.Size() % 2 != 0 {
			return false
		}
		r, _ := subtrees.Pop()
		l,  _ := subtrees.Pop()

		right, left := r.(*TreeNode), l.(*TreeNode)

		if right == nil || left == nil {
			if left != right {
				return false
			}
			continue
		}

		if right.Val != left.Val {
			return false
		}

		subtrees.Push(left.Left)
		subtrees.Push(right.Right)

		subtrees.Push(left.Right)
		subtrees.Push(right.Left)
	}

	return true
}
