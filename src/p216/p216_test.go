package p216

import (
	"testing"
	. "testing/utils"
)

func TestCombinationSum3(t *testing.T) {
	testCase := func(in0, in1 int, want [][]int) {
		got := combinationSum3(in0, in1)
		if !Is2DIntsEqualWithoutOrder(got, want) {
			t.Errorf("Case %v, %v: expected %v, got %v",
				in0, in1, want, got)
		}
	}

	testCase(3, 7, [][]int{{1, 2, 4}})
	testCase(3, 9, [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}})
}
