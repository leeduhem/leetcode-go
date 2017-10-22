package p059

import "testing"
import . "testing/utils"

func testCase(in0 int, out0 [][]int, t *testing.T) {
	out1 := generateMatrix(in0)
	if ! Is2DIntMatrixEqual(out0, out1) {
		t.Errorf("Case %v: expected %v, got %v", in0, out0, out1)
	}
}

func TestGenerateMatrix(t *testing.T) {
	in0  := 3
	out0 := [][]int {
		{1,2,3},
		{8,9,4},
		{7,6,5},
	}
	testCase(in0, out0, t)

	in1  := 1
	out1 := [][]int { {1} }
	testCase(in1, out1, t)
}
