package p092

import (
	"testing"
)

func testCase(in0 *ListNode, in1, in2 int, want *ListNode, t *testing.T) {
	oin0 := deepCopyList(in0)
	got := reverseBetween(in0, in1, in2)
	if ! isListEqual(got, want) {
		t.Errorf("Case %v, %v, %v: expected %v, got %v", oin0, in1, in2, want, got)
	}
}

func TestReverseBetween(t *testing.T) {
	in00 := ints2List([]int {1,2,3,4,5})
	in01, in02 := 2, 4
	want0 := ints2List([]int {1,4,3,2,5})
	testCase(in00, in01, in02, want0, t)

	in10 := ints2List([]int {5})
	in11, in12 := 1,1
	want1 := ints2List([]int {5})
	testCase(in10, in11, in12, want1, t)

	in20 := ints2List([]int {3,5})
	in21, in22 := 1,2
	want2 := ints2List([]int {5,3})
	testCase(in20, in21, in22, want2, t)
}
