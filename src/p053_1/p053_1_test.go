package p053_1

import "testing"

func testCase(in0 []int, out0 int, t *testing.T) {
	out1 := maxSubArray(in0)
	if out0 != out1 {
		t.Errorf("Case %v: expected %v, got %v", in0, out0, out1)
	}
}

func TestMaxSubArray(t *testing.T) {
	in0  := []int {-2,1,-3,4,-1,2,1,-5,4}
	out0 := 6
	testCase(in0, out0, t)

	in1  := []int {1}
	out1 := 1
	testCase(in1, out1, t)

	in2  := []int {-2, 1}
	out2 := 1
	testCase(in2, out2, t)

	in3  := []int {-1, -2}
	out3 := -1
	testCase(in3, out3, t)
}
