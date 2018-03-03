package p152

import (
	"testing"
)

func testCase(in []int, want int, t *testing.T) {
	got := maxProduct(in)
	if got != want {
		t.Errorf("Case %v: expected %v, got %v", in, want, got)
	}
}

func TestMaxProduct(t *testing.T) {
	in0 := []int {2,3,-2,4}
	want0 := 6
	testCase(in0, want0, t)
}
