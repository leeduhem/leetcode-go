package p165

import (
	"testing"
)

func testCase(in0, in1 string, want int, t *testing.T) {
	got := compareVersion(in0, in1)
	if got != want {
		t.Errorf("Case %v %v: expected %v, got %v", in0, in1, want, got)
	}
}

func TestCompareVersion(t *testing.T) {
	in00, in01 := "0.1", "1.1"
	want0 := -1
	testCase(in00, in01, want0, t)

	in10, in11 := "13.37", "1.2"
	want1 := 1
	testCase(in10, in11, want1, t)

	in20, in21 := "01", "1"
	want2 := 0
	testCase(in20, in21, want2, t)

	in30, in31 := "1.0", "1"
	want3 := 0
	testCase(in30, in31, want3, t)
}
