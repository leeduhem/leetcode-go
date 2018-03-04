package p160

import (
	"testing"
)

func testCase(in0, in1, want *ListNode, t *testing.T) {
	got := getIntersectionNode(in0, in1)
	if got != want {
		t.Errorf("Case %v, %v: expected %v, got %v", in0, in1, want, got)
	}
}

func TestGetIntersectionNode(t *testing.T) {
	in00 := ints2List([]int {1, 2, 3})
	in01 := &ListNode{1, &ListNode{2, in00}}
	in02 := &ListNode{1, &ListNode{2, &ListNode{3, in00}}}
	want0 := in00
	testCase(in01, in02, want0, t)
}
