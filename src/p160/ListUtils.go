package p160

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func (l *ListNode) String() string {
	d := 0
	str := ""

	for n := l; n != nil; n = n.Next {
		str += fmt.Sprintf("&{%v -> ", n.Val)
		d++
	}
	str += "<nil>"
	str += strings.Repeat("}", d)

	return str
}

func isListEqual(l1, l2 *ListNode) bool {
	for ; l1 != nil && l2 != nil; l1, l2 = l1.Next, l2.Next {
		if l1.Val != l2.Val {
			return false
		}
	}

	if l1 != nil || l2 != nil {
		return false
	}

	return true
}

func deepCopyList(l *ListNode) *ListNode{
	l1 := (*ListNode)(nil)

	for tail := &l1; l != nil; l = l.Next {
		*tail = &ListNode{ l.Val, nil }
		tail = &(*tail).Next
	}

	return l1
}

func ints2List(ns []int) *ListNode {
	head := (*ListNode)(nil)
	tail := &head

	for _, n := range ns {
		*tail = &ListNode{ n, nil }
		tail  = &(*tail).Next
	}
	return head
}
