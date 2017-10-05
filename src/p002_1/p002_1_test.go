package p002_1

import "testing"

func testCase(a1,a2,s0 *ListNode, t *testing.T) {
	s1 := addTwoNumbers(a1, a2)

	for ; s0 != nil && s1 != nil; s0,s1 = s0.Next, s1.Next {
		if (s0.Val != s1.Val) {
			t.Fatalf("Case %v %v: expected %v, got %v", a1, a2, s0, s1)
		}
	}

	if s0 != nil || s1 != nil {
		t.Fatalf("Case %v %v: expected %v, got %v", a1, a2, s0, s1)
	}
}

func TestAddTwoNumbers(t *testing.T) {
	l11 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l12 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	r10 := &ListNode{7, &ListNode{0, &ListNode{8, nil}}}
	testCase(l11, l12, r10, t)


	l21 := &ListNode{5, nil}
	r20 := &ListNode{0, &ListNode{1, nil}}
	testCase(l21, l21, r20, t)
}
