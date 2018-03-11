package p169_1

import (
	"testing"
)

func TestMajorityElement(t *testing.T) {
	testCase := func(in []int, want int) {
		got := majorityElement(in)
		if got != want {
			t.Errorf("Case %v: expected %v, got %v", in, want, got)
		}
	}

	testCase([]int{1,1,2}, 1)
	testCase([]int{6,5,5}, 5)
}
