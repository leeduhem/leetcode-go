package p054

import "testing"
import . "testing/utils"

func testCase(in0 [][]int, out0 []int, t *testing.T) {
	out1 := spiralOrder(in0)
	if ! IsIntsEqual(out0, out1) {
		t.Errorf("Case %v: expected %v, got %v", in0, out0, out1)
	}
}

func TestSpiralOrder(t *testing.T) {
	in0  := [][]int {
		{1,2,3,},
		{4,5,6,},
		{7,8,9,},
	}
	out0 := []int {1,2,3,6,9,8,7,4,5}
	testCase(in0, out0, t)

	in1  := [][]int {}
	out1 := []int {}
	testCase(in1, out1, t)

	in2  := [][]int {
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9,10,11,12},
	}
	out2 := []int {1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}
	testCase(in2, out2, t)

	in3  := [][]int {
		{3},
		{2},
	}
	out3 := []int {3, 2}
	testCase(in3, out3, t)
}
