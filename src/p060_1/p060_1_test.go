package p060_1

import "testing"

func testCase(in0, in1 int, out0 string, t *testing.T) {
	out1 := getPermutation(in0, in1)
	if out0 != out1 {
		t.Errorf("Case %v %v: expected %v, got %v", in0, in1, out0, out1)
	}
}

func TestGetPermutation(t *testing.T) {
	in00 := 3
	in01 := 2
	out0 := "132"
	testCase(in00, in01, out0, t)

	in10 := 4
	in11 := 14
	out1 := "3142"
	testCase(in10, in11, out1, t)
}
