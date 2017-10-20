package p048

import "testing"
import . "testing/utils"

func testCase(in0, out0 [][]int, t *testing.T) {
	in01 := DeepCopy2DMatrix(in0)
	rotate(in0)
	if ! Is2DIntMatrixEqual(out0, in0) {
		t.Errorf("Case \n%v\n: expected \n%v, got \n%v", IntMatrix(in01), IntMatrix(out0), IntMatrix(in0))
	}
}

func TestRotate(t *testing.T) {
	in1  := [][]int {
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	out1 := [][]int {
		{7,4,1},
		{8,5,2},
		{9,6,3},
	}
	testCase(in1, out1, t)

	in2  := [][]int {
		{ 5, 1, 9,11},
		{ 2, 4, 8,10},
		{13, 3, 6, 7},
		{15,14,12,16},
	}
	out2 := [][]int {
		{15,13, 2, 5},
		{14, 3, 4, 1},
		{12, 6, 8, 9},
		{16, 7,10,11},
	}
	testCase(in2, out2, t)
}
