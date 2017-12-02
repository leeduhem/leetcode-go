package p123

import (
	"testing"
)

func testCase(in []int, want int, t *testing.T) {
	got := maxProfit(in)
	if got != want {
		t.Errorf("Case %v: expected %v, got %v", in, want, got)
	}
}

func TestMaxProfit(t *testing.T) {
	in0 := []int { 7, 1, 5, 3, 6, 4 }
	want0 := 7
	testCase(in0, want0, t)

	in1 := []int { 7, 6, 4, 3, 1 }
	want1 := 0
	testCase(in1, want1, t)
}
