package p188

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	testCase := func(in0 int, in1 []int, want int) {
		got := maxProfit(in0, in1)
		if got != want {
			t.Errorf("Case %v, %v: expected %v, got %v", in0, in1, want, got)
		}
	}

	testCase(2, []int{}, 0)
}
