package p063

import (
	"testing"
)

func testCase(in0 [][]int, out0 int, t *testing.T) {
	out1 := uniquePathsWithObstacles(in0)
	if out0 != out1 {
		t.Errorf("Case %v: expected %v, got %v", in0, out0, out1)
	}
}

func TestUniquePathsWithObstacles(t *testing.T) {
	in0  := [][]int {
		{0,0,0},
		{0,1,0},
		{0,0,0},
	}
	out0 := 2
	testCase(in0, out0, t)
}
