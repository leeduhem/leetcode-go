package p174

import (
	"testing"
	. "testing/utils"
)

func TestCalculateMinimumHP(t *testing.T) {
	testCase := func(in [][]int, want int) {
		in0 := DeepCopy2DIntMatrix(in)
		got := calculateMinimumHP(in)
		if got != want {
			t.Errorf("Case %v: expected %v, got %v", in0, want, got)
		}
	}

	in0 := [][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5},
	}
	testCase(in0, 7)

	in1 := [][]int{{0}}
	testCase(in1, 1)

	in2 := [][]int{{-3, 5}}
	testCase(in2, 4)

	in3 := [][]int{{0, 5}, {-2, -3}}
	testCase(in3, 1)
}
