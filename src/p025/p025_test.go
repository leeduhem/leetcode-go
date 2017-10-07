package p025

import (
	"testing"
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

func array2list(a []int) *ListNode {
	head := &ListNode{0, nil}
	tail := head

	for _,v := range a {
		tail.Next = &ListNode{v, nil}
		tail = tail.Next
	}

	return head.Next
}


func testCase(in0 *ListNode, in1 int, out0 *ListNode, t *testing.T) {
	out1 := reverseKGroup(in0, in1)

	i0, i1 := in0, in1
	o0, o1 := out0, out1

	for ; out0 != nil && out1 != nil; out0, out1 = out0.Next, out1.Next {
		if out0.Val != out1.Val {
			t.Fail()
		}
	}

	if out0 != nil || out1 != nil {
		t.Fail()
	}

	if t.Failed() {
		t.Errorf("Case %v %v: expected %v, got %v", i0, i1, o0, o1)
	}
}

func TestReverseKGroup(t *testing.T) {
	in10  := array2list([]int{1, 2, 3, 4, 5})
	in11  := 2
	out10 := array2list([]int{2, 1, 4, 3, 5})
	testCase(in10, in11, out10, t)

	in20  := array2list([]int{1, 2, 3, 4, 5})
	in21  := 3
	out20 := array2list([]int{3, 2, 1, 4, 5})
	testCase(in20, in21, out20, t)
}
