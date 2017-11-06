package p097

import (
	"testing"
)

func testCase(in0, in1, in2 string, want bool, t *testing.T) {
	got := isInterleave(in0, in1, in2)
	if got != want {
		t.Errorf("Case %v, %v, %v: expected %v, got %v", in0, in1, in2, want, got)
	}
}

func TestIsInterleave(t *testing.T) {
	in00, in01, in02 := "aabcc", "dbbca", "aadbbcbcac"
	want0 := true
	testCase(in00, in01, in02, want0, t)

	in10, in11, in12 := "aabcc", "dbbca", "aadbbbaccc"
	want1 := false
	testCase(in10, in11, in12, want1, t)

	in20, in21, in22 := "a", "", "a"
	want2 := true
	testCase(in20, in21, in22, want2, t)

	in30, in31, in32 := "aabccabc", "dbbabc", "aabdbbccababcc"
	want3 := true
	testCase(in30, in31, in32, want3, t)
}
