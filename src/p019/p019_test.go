package p019

import "testing"

func testCase(in0 *ListNode, in1 int, out0 *ListNode, t *testing.T) {
	out1 := removeNthFromEnd(in0, in1)

	for ; out0 != nil && out1 != nil; out0, out1 = out0.Next, out1.Next {
		if out0.Val != out1.Val {
			t.Fatalf("Case %v %v: expected %v, got %v", in0, in1, out0, out1)
		}
	}

	if out0 != nil || out1 != nil {
		t.Fatalf("Case %v %v: expected %v, got %v", in0, in1, out0, out1)
	}
}

func TestRemoveNthFromEnd(t *testing.T) {
	in10 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	in11 := 2
	out10 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{5, nil}}}}
	testCase(in10, in11, out10, t)

	in20 := &ListNode{1, nil}
	in21 := 1
	out20 := (*ListNode)(nil)
	testCase(in20, in21, out20, t)
}
