package p073_1

import (
	"testing"
	. "testing/utils"
)


func testCase(in0, want [][]int, t *testing.T) {
	oin0 := DeepCopy2DIntMatrix(in0)

	setZeroes(oin0)
	if ! Is2DIntMatrixEqual(oin0, want) {
		t.Errorf("Case %v: expected %v, got %v", in0, want, oin0)
	}
}

func TestSetZeroes(t *testing.T) {
	in0 := [][]int {
		{1,2,3},
		{4,0,6},
		{7,8,9},
	}
	want0 := [][]int {
		{1,0,3},
		{0,0,0},
		{7,0,9},
	}
	testCase(in0, want0, t)
}
