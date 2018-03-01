package p143

import (
	"testing"
)

func testCase(in, want *ListNode, t *testing.T) {
	in0 := deepCopyList(in)
	reorderList(in)
	if isListEqual(in, in0) {
		t.Errorf("case %v: expected %v, got %v", in0, want, in)
	}
}

func TestReorderList(t *testing.T) {
	in0 := ints2List([]int {1,2,3,4})
	want0 := ints2List([]int {1,4,2,3})
	testCase(in0, want0, t)
}
