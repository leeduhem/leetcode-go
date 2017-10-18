package p043

import "testing"

func testCase(in0, in1, out0 string, t *testing.T) {
	out1 := multiply(in0, in1)
	if out0 != out1 {
		t.Errorf("Case %v %v: expected %v, got %v", in0, in1, out0, out1)
	}
}

func TestMultiply(t *testing.T) {
	in10  := "456"
	in11  := "789"
	out10 := "359784"
	testCase(in10, in11, out10, t)

	in20  := "9133"
	in21  := "0"
	out20 := "0"
	testCase(in20, in21, out20, t)
}
