package p024

import "testing"

type ListNode struct {
    Val int
    Next *ListNode
}

func testCase(in0, out0 *ListNode, t *testing.T) {
	out1 := swapPairs(in0)

	for ; out0 != nil && out1 != nil; out0, out1 = out0.Next, out1.Next {
		if out0.Val != out1.Val {
			t.Fatalf("Case %v: expected %v, got %v", in0, out0, out1)
		}
	}

	if out0 != nil || out1 != nil {
		t.Fatalf("Case %v: expected %v, got %v", in0, out0, out1)
	}
}

func TestSwapPairs(t *testing.T) {
	in0  := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	out0 := &ListNode{2, &ListNode{1, &ListNode{4, &ListNode{3, nil}}}}
	testCase(in0, out0, t)
}
