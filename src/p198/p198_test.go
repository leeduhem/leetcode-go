package p198

import (
	"testing"
)

func TestRob(t *testing.T) {
	testCase := func(in []int, want int) {
		got := rob(in)
		if got != want {
			t.Errorf("Case %v: expected %v, got %v", in, want, got)
		}
	}

	testCase([]int{}, 0)
	testCase([]int{1}, 1)
	testCase([]int{1,3,1}, 3)
	testCase([]int{2,1,1,2}, 4)
}
