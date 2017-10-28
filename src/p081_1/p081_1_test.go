package p081_1

import (
	"testing"
)

func testCase(in0 []int, in1 int, want bool, t *testing.T) {
	got := search(in0, in1)
	if got != want {
		t.Errorf("Case %v, %v: expected %v, got %v", in0, in1, want, got)
	}
}

func TestSearch(t *testing.T) {
	in00 := []int {4,5,6,7,0,1,2}
	in01 := 7
	want0 := true
	testCase(in00, in01, want0, t)

	in10 := []int {}
	in11 := 5
	want1 := false
	testCase(in10, in11, want1, t)

	in20 := []int {1}
	in21 := 0
	want2 := false
	testCase(in20, in21, want2, t)

	in30 := []int {1,1}
	in31 := 0
	want3 := false
	testCase(in30, in31, want3, t)

	in40 := []int {1,3}
	in41 := 3
	want4 := true
	testCase(in40, in41, want4, t)

	in50 := []int {1,3}
	in51 := 2
	want5 := false
	testCase(in50, in51, want5, t)

	in60 := []int {3,1}
	in61 := 1
	want6 := true
	testCase(in60, in61, want6, t)
}
