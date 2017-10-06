package p021

import "testing"

func testCase(in0, in1, out0 *ListNode, t *testing.T) {
	out1 := mergeTwoLists(in0, in1)

	for ; out0 != nil && out1 != nil; out0, out1 = out0.Next, out1.Next {
		if out0.Val != out1.Val {
			t.Errorf("Case %v %v: expected %v, got %v", in0, in1, out0, out1)
		}
	}
}

func TestMergeTwoLists(t *testing.T) {
	in10 := &ListNode{1, &ListNode{3, nil}}
	in11 := &ListNode{2, &ListNode{4, nil}}
	out10 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	testCase(in10, in11, out10, t)

	in20 := (*ListNode)(nil)
	in21 := &ListNode{0, nil}
	out20 := &ListNode{0, nil}
	testCase(in20, in21, out20, t)
}
