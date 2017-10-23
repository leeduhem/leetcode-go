package p061

import "testing"

type ListNode struct {
    Val int
    Next *ListNode
}

func IsListEqual(l, l1 *ListNode) bool {
	for ; l != nil && l1 != nil; l, l1 = l.Next, l1.Next {
		if l.Val != l1.Val {
			return false
		}
	}

	if l != nil || l1 != nil {
		return false
	}

	return true
}

func testCase(in0 *ListNode, in1 int, out0 *ListNode, t *testing.T) {
	out1 := rotateRight(in0, in1)
	if ! IsListEqual(out0, out1) {
		t.Errorf("Case %v %v: expected %v, got %v", in0, in1, out0, out1)
	}
}

func TestRotateRight(t *testing.T) {
	in00 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	in01 := 2
	out0 := &ListNode{4, &ListNode{5, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}}}
	testCase(in00, in01, out0, t)

	in10 := &ListNode{}
	in11 := 0
	out1 := &ListNode{}
	testCase(in10, in11, out1, t)

	in20 := &ListNode{1, nil}
	in21 := 1
	out2 := &ListNode{1, nil}
	testCase(in20, in21, out2, t)

	in30 := &ListNode{1, &ListNode{2, nil}}
	in31 := 3
	out3 := &ListNode{2, &ListNode{1, nil}}
	testCase(in30, in31, out3, t)

	in40 := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	in41 := 2000000000
	out4 := &ListNode{2, &ListNode{3, &ListNode{1, nil}}}
	testCase(in40, in41, out4, t)
}
