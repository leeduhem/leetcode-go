package p083

import (
	"testing"
)

func testCase(in0, want *ListNode, t *testing.T) {
	oin0 := deepCopyList(in0)
	got := deleteDuplicates(in0)
	if ! isListEqual(got, want) {
		t.Errorf("Case %v: expected %v, got %v", oin0, want, got)
	}
}

func TestDeleteDuplicates(t *testing.T) {
	in0   := ints2List([]int {1,1,2})
	want0 := ints2List([]int {1,2})
	testCase(in0, want0, t)

	in1   := ints2List([]int {1,1,2,3,3})
	want1 := ints2List([]int {1,2,3})
	testCase(in1, want1, t)
}
