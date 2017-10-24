package p067

import (
	"testing"
)

func testCase(in0, in1, out0 string, t *testing.T) {
	out1 := addBinary(in0, in1)
	if out0 != out1 {
		t.Errorf("Case %v %v: expected %v, got %v", in0, in1, out0, out1)
	}
}

func TestAddBinary(t *testing.T) {
	in00 := "11"
	in01 := "1"
	out0 := "100"
	testCase(in00, in01, out0, t)

	in10 := "100"
	in11 := "110010"
	out1 := "110110"
	testCase(in10, in11, out1, t)
}
