package p134

import (
	"testing"
)

func testCase(in0, in1 []int, want int, t *testing.T) {
	got := canCompleteCircuit(in0, in1)
	if got != want {
		t.Errorf("Case %v, %v: expected %v, got %v", in0, in1, want, got)
	}
}

func TestCanCompleteCircuit(t *testing.T) {
	in00 := []int {5}
	in01 := []int {4}
	want0 := 0
	testCase(in00, in01, want0, t)

	in10 := []int {4}
	in11 := []int {5}
	want1 := -1
	testCase(in10, in11, want1, t)
}
