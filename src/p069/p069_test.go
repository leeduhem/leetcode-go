package p069

import (
	"testing"
)

func testCase(in0, out0 int, t *testing.T) {
	out1 := mySqrt(in0)
	if out0 != out1 {
		t.Errorf("Case %v: expected %v, got %v", in0, out0, out1)
	}
}

func TestMySqrt(t *testing.T) {
	in0  := 4
	out0 := 2
	testCase(in0, out0, t)
}
