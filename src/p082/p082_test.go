package p082

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

func deepCopyList(l *ListNode) *ListNode{
	l1 := (*ListNode)(nil)

	for tail := &l1; l != nil; l = l.Next {
		*tail = &ListNode{ l.Val, nil }
		tail = &(*tail).Next
	}

	return l1
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

func testCase(in0, want *ListNode, t *testing.T) {
	oin0 := deepCopyList(in0)
	got := deleteDuplicates(in0)
	if ! isListEqual(got, want) {
		t.Errorf("Case %v: expected %v, got %v", oin0, want, got)
	}
}

func TestDeleteDuplicates(t *testing.T) {
	in0 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, nil}}}}}}}
	want0 := &ListNode{1, &ListNode{2, &ListNode{5, nil}}}
	testCase(in0, want0, t)

	in1 := &ListNode{1, &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}}}
	want1 := &ListNode{2, &ListNode{3, nil}}
	testCase(in1, want1, t)

	in2 := (*ListNode)(nil)
	want2 := (*ListNode)(nil)
	testCase(in2, want2, t)
}
