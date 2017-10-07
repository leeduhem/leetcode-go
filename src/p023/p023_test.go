package p023

import "testing"

func testCase(in0 []*ListNode, out0 *ListNode, t *testing.T) {
	out1 := mergeKLists(in0)

	for ; out0 != nil && out1 != nil; out0, out1 = out0.Next, out1.Next {
		if out0.Val != out1.Val {
			t.Errorf("Case %v: expected %v, got %v", in0, out0, out1)
		}
	}

	if out0 != nil || out1 != nil {
		t.Errorf("Case %v: expected %v, got %v", in0, out0, out1)
	}
}

func TestMergeTwoLists(t *testing.T) {
	in10 := []*ListNode{
		&ListNode{1, &ListNode{3, nil}},
		&ListNode{2, &ListNode{4, nil}},
	}
	out10 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	testCase(in10, out10, t)

	in20 := []*ListNode{
		(*ListNode)(nil),
		&ListNode{0, nil},
	}
	out20 := &ListNode{0, nil}
	testCase(in20, out20, t)

	in30 := []*ListNode{
		&ListNode{1, &ListNode{2, &ListNode{3, nil}}},
		&ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, nil}}}},
	}
	out30 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, nil}}}}}}}
	testCase(in30, out30, t)
}
