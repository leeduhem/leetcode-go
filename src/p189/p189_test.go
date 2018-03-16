package p189

import (
	"testing"
	. "testing/utils"
)

func TestRotate(t *testing.T) {
	testCase := func(in0 []int, in1 int, want []int) {
		got := make([]int, len(in0))
		copy(got, in0)
		rotate(got, in1)
		if ! IsIntsEqual(got, want) {
			t.Errorf("Case %v, %v: expected %v, got %v", in0, in1, want, got)
		}
	}

	testCase([]int{1,2,3,4,5,6,7}, 3, []int{5,6,7,1,2,3,4})
}
