package p002_1

import "testing"

func checkResult(r,s *ListNode, t *testing.T) {
	for ; s != nil && r != nil; s,r = s.Next, r.Next {
		if (s.Val != r.Val) {
			t.Errorf("Expected %c, got %c", r.Val, s.Val)
		}
	}

	if (r != nil) {
		t.Errorf("Expected %c", r.Val)
	}
	if (s != nil) {
		t.Errorf("Got %c", s.Val)
	}
}

func TestAddTwoNumbers(t *testing.T) {
	l11 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l12 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	r1  := &ListNode{7, &ListNode{0, &ListNode{8, nil}}}
	s1  := addTwoNumbers(l11, l12)
	checkResult(r1, s1, t)


	l21 := &ListNode{5, nil}
	r2  := &ListNode{0, &ListNode{1, nil}}
	s2  := addTwoNumbers(l21, l21)
	checkResult(r2, s2, t)
}
