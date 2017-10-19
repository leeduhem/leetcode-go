package p046

import "testing"
import . "testing/utils"

func testCase(in0 []int, out0 [][]int, t *testing.T) {
	out1 := permute(in0)
	if ! CmpMatrix(out0, out1) {
		t.Errorf("Case %v: expected %v, got %v", in0, out0, out1)
	}
}

func TestPermute(t *testing.T) {
	in1  := []int {1,2,3}
	out1 := [][]int {
		{1,2,3},
		{1,3,2},
		{2,1,3},
		{2,3,1},
		{3,1,2},
		{3,2,1},
	}
	testCase(in1, out1, t)
}
