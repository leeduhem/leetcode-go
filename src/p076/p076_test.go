package p076

import (
	"testing"
)

func testCase(in0, in1, want string, t *testing.T) {
	got := minWindow(in0, in1)
	if got != want {
		t.Errorf("Case %v, %v: expected %v, got %v", in0, in1, want, got)
	}
}

func TestMinWindow(t *testing.T) {
	in00 := "ADOBECODEBANC"
	in01 := "ABC"
	want0 := "BANC"
	testCase(in00, in01, want0, t)

	in10  := "a"
	in11  := "a"
	want1 := "a"
	testCase(in10, in11, want1, t)

	in20  := "bba"
	in21  := "ab"
	want2 := "ba"
	testCase(in20, in21, want2, t)

	in30  := "aa"
	in31  := "aa"
	want3 := "aa"
	testCase(in30, in31, want3, t)

	in40  := "bdab"
	in41  := "ab"
	want4 := "ab"
	testCase(in40, in41, want4, t)
}
