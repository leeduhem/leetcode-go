package p087

import (
	"testing"
)

func testCase(in0, in1 string, want bool, t *testing.T) {
	got := isScramble(in0, in1)
	if got != want {
		t.Errorf("Case %v, %v: expected %v, got %v", in0, in1, want, got)
	}
}

func TestIsScramble(t *testing.T) {
	in00 := "great"
	in01 := "rgeat"
	want0 := true
	testCase(in00, in01, want0, t)

	in10 := "great"
	in11 := "rgtae"
	want1 := true
	testCase(in10, in11, want1, t)

	in20 := "aa"
	in21 := "ab"
	want2 := false
	testCase(in20, in21, want2, t)

	in30 := "abcd"
	in31 := "bdac"
	want3 := false
	testCase(in30, in31, want3, t)
}
