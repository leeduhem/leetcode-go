package p050_1

import (
	"math"
	"testing"
)

func testCase(in0 float64, in1 int, out0 float64, t *testing.T) {
	out1 := myPow(in0, in1)

	if math.Abs(out1 - out0) > 5e-6 {
		t.Logf("XXX: %v", math.Abs(out1 - out0))
		t.Errorf("Case %v %v: expected %v, got %v", in0, in1, out0, out1)
	}
}

func TestMyPow(t *testing.T) {
	in10  := 3.0
	in11  := 2
	out10 := 9.0
	testCase(in10, in11, out10, t)

	in20  := 34.00515
	in21  := -3
	out20 := 0.00003
	testCase(in20, in21, out20, t)

	in30  := 0.00001
	in31  := 2147483647
	out30 := 0.0
	testCase(in30, in31, out30, t)

	in40  := 8.95371
	in41  := -1
	out40 := 0.11169
	testCase(in40, in41, out40, t)

	in50  := 42.38803
	in51  := 1
	out50 := 42.38803
	testCase(in50, in51, out50, t)
}
