package p062

import "testing"

func testCase(in0, in1, out0 int, t *testing.T) {
	out1 := uniquePaths(in0, in1)
	if out0 != out1 {
		t.Errorf("Case %v %v: expected %v, got %v", in0, in1, out0, out1)
	}
}

func TestUniquePaths(t *testing.T) {
	in00 := 3
	in01 := 2
	out0 := 3
	testCase(in00, in01, out0, t)

	in10 := 1
	in11 := 1
	out1 := 1
	testCase(in10, in11, out1, t)
}
