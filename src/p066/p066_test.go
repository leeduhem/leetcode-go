package p066

import (
	"testing"
)

import . "testing/utils"

func testCase(in0, out0 []int, t *testing.T) {
	in0o := make([]int, len(in0))
	copy(in0o, in0)

	out1 := plusOne(in0)
	if ! IsIntsEqual(out0, out1) {
		t.Errorf("Case %v: expected %v, got %v", in0o, out0, out1)
	}
}

func TestPlusOne(t *testing.T) {
	in0  := []int {1,2,3}
	out0 := []int {1,2,4}
	testCase(in0, out0, t)

	in1  := []int {9}
	out1 := []int {1,0}
	testCase(in1, out1, t)
}
