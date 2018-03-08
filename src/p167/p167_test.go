package p167

import (
	"testing"
	. "testing/utils"
)

func TestTwoSum(t *testing.T) {
	testCase := func(in0 []int, in1 int, want []int) {
		got := twoSum(in0, in1)
		if ! IsIntsEqual(got, want) {
			t.Errorf("Case %v, %v: expected %v, got %v", in0, in1, want, got)
		}
	}

	testCase([]int{2, 7, 11, 15}, 9, []int{1, 2})
	testCase([]int{3,24,50,79,88,150,345}, 200, []int{3,6})
}
