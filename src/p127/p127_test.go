package p127

import (
	"testing"
)

func testCase(in0, in1 string, in2 []string, want int, t *testing.T) {
	got := ladderLength(in0, in1, in2)
	if got != want {
		t.Errorf("Case %v, %v, %v: expected %v, got %v", in0, in1, in2, want, got)
	}
}

func TestLadderLength(t *testing.T) {
	in00, in01 := "hit", "cog"
	in02 := []string { "hot","dot","dog","lot","log","cog" }
	want0 := 5
	testCase(in00, in01, in02, want0, t)

	in10, in11 := "hit", "cog"
	in12 := []string { "hot","dot","dog","lot","log" }
	want1 := 0
	testCase(in10, in11, in12, want1, t)
}
