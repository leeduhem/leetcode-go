package p070

import (
	"testing"
)

func testCase(in0, out0 int, t *testing.T) {
	out1 := climbStairs(in0)
	if out0 != out1 {
		t.Errorf("Case %v: expected %v, got %v", in0, out0, out1)
	}
}

func TestClimbStairs(t *testing.T) {
	in0  := 1
	out0 := 1
	testCase(in0, out0, t)

	in1  := 2
	out1 := 2
	testCase(in1, out1, t)

	in2  := 3
	out2 := 3
	testCase(in2, out2, t)
}
