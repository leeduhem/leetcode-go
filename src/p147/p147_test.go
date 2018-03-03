package p147

import (
	"testing"
)

func testCase(in, want *ListNode, t *testing.T) {
	in0 := deepCopyList(in)
	got := insertionSortList(in)
	if ! isListEqual(got, want) {
		t.Errorf("Case %v: expected %v, got %v", in0, want, got)
	}
}

func TestInsertionSortList(t *testing.T) {
	in0 := ints2List([]int {1,3,2,4})
	want0 := ints2List([]int {1,2,3,4})
	testCase(in0, want0, t)
}
