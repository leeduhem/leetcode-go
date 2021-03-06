package p047

import "testing"
import . "testing/utils"

func testCase(in0 []int, out0 [][]int, t *testing.T) {
	out1 := permuteUnique(in0)
	if ! Is2DIntsEqualWithoutOrder(out0, out1) {
		t.Errorf("Case %v: expected %v, got %v", in0, out0, out1)
	}
}

func TestPermuteUnique(t *testing.T) {
	in1  := []int {1,1,2}
	out1 := [][]int {
		{1,1,2},
		{1,2,1},
		{2,1,1},
	}
	testCase(in1, out1, t)
}
