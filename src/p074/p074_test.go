package p074

import (
	"testing"
)

func testCase(in0 [][]int, in1 int, want bool, t *testing.T) {
	got := searchMatrix(in0, in1)
	if got != want {
		t.Errorf("Case %v, %v: expected %v, got %v", in0, in1, want, got)
	}
}

func TestSearchMatrix(t *testing.T) {
	in00 := [][]int {
		{1,   3,  5,  7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	in01 := 3
	want0 := true
	testCase(in00, in01, want0, t)

	in10 := [][]int {
		{1,3},
	}
	in11 := 2
	want1 := false
	testCase(in10, in11, want1, t)

	in20 := [][]int {
		{1},
		{3},
	}
	in21 := 3
	want2 := true
	testCase(in20, in21, want2, t)
}
