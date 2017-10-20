package p050

import "testing"

func testCase(in0 float64, in1 int, out0 float64, t *testing.T) {
	out1 := myPow(in0, in1)

	// Should NOT compare floating point numbers directly.
	if out0 != out1 {
		t.Errorf("Case %v %v: expected %v, got %v", in0, in1, out0, out1)
	}
}

func TestMyPow(t *testing.T) {
	in10  := 3.0
	in11  := 2
	out10 := 9.0
	testCase(in10, in11, out10, t)
}
