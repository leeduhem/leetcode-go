package p064

import (
	"testing"
)

func testCase(in0 [][]int, out0 int, t *testing.T) {
	out1 := minPathSum(in0)
	if out0 != out1 {
		t.Errorf("Case %v: expected %v, got %v", in0, out0, out1)
	}
}

func TestMinPathSum(t *testing.T) {
	in0  := [][]int {
		{1,3,1},
		{1,5,1},
		{4,2,1},
	}
	out0 := 7
	testCase(in0, out0, t)

	in1  := [][]int {
		{1,2,5},
		{3,2,1},
	}
	out1 := 6
	testCase(in1, out1, t)

	in2  := [][]int {
		{1,2},
		{1,1},
	}
	out2 := 3
	testCase(in2, out2, t)
}
