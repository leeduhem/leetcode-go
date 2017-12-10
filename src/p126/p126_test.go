package p126

import (
	"testing"
	. "testing/utils"
)

func testCase(in0, in1 string, in2 []string, want [][]string, t *testing.T) {
	got := findLadders(in0, in1, in2)
	if len(got) != len(want) {
		t.Errorf("Case %v, %v, %v: expected %v, got %v", in0, in1, in2, want, got)
	}

outerLoop:
	for _, ss1 := range got {
		for _, ss2 := range want {
			if IsStrsEqual(ss1, ss2) {
				continue outerLoop
			}
		}
		t.Errorf("Case %v, %v, %v: expected %v, got %v", in0, in1, in2, want, got)
	}
}

func TestFindLadders(t *testing.T) {
	in00, in01 := "hit", "cog"
	in02 := []string { "hot","dot","dog","lot","log","cog" }
	want0 := [][]string {
		{ "hit","hot","dot","dog","cog" },
		{ "hit","hot","lot","log","cog" },
	}
	testCase(in00, in01, in02, want0, t)

	in10, in11 := "a", "c"
	in12 := []string { "a","b","c" }
	want1 := [][]string { {"a", "c"} }
	testCase(in10, in11, in12, want1, t)
}
