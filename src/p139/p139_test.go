package p139

import (
	"testing"
)

func testCase(in0 string, in1 []string, want bool, t *testing.T) {
	got := wordBreak(in0, in1)
	if got != want {
		t.Errorf("Case %v %v: expected %v, got %v", in0, in1, want, got)
	}
}

func TestWordBreak(t *testing.T) {
	in00 := "leetcode"
	in01 := []string { "leet", "code" }
	want0 := true
	testCase(in00, in01, want0, t)

	in10 := "cars"
	in11 := []string { "car","ca","rs" }
	want1 := true
	testCase(in10, in11, want1, t)
}
