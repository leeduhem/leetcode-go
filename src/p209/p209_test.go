package p209

import (
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	testCase := func(in0 int, in1 []int, want int) {
		got := minSubArrayLen(in0, in1)
		if got != want {
			t.Errorf("Case %v, %v: expected %v, got %v", in0, in1, want, got)
		}
	}

	testCase(7, []int{2, 3, 1, 2, 4, 3}, 2)
}
