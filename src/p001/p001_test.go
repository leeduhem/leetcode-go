package p001

import (
	"testing"
	. "testing/utils"
)

func TestTwoSum(t *testing.T) {
	testCase := func(in0 []int, in1 int, want []int) {
		got := twoSum(in0, in1)
		if !IsIntsEqual(got, want) {
			t.Errorf("Case %v, %v: expected %v, got %v",
				in0, in1, want, got)
		}
	}
	testCase([]int{3, 2, 4}, 6, []int{1, 2})
}
