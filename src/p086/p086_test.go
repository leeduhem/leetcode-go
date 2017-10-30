package p086

import (
	"testing"
)

func testCase(in0 *ListNode, in1 int, want *ListNode, t *testing.T) {
	oin0 := deepCopyList(in0)
	got := partition(in0, in1)
	if ! doesListEqual(got, want) {
		t.Errorf("Case %v, %v: expected %v, got %v", oin0, in1, want, got)
	}
}

func TestPartition(t *testing.T) {
	in00 := ints2List([]int {1,4,3,2,5,2})
	in01 := 3
	want0 := ints2List([]int {1,2,2,4,3,5})
	testCase(in00, in01, want0, t)
}
