package p135

import (
	"testing"
)

func testCase(in []int, want int, t *testing.T) {
	got := candy(in)
	if got != want {
		t.Errorf("Case %v: expected %v, got %v", in, want, got)
	}
}

func TestCandy(t *testing.T) {
	in0 := []int {0}
	want0 := 1
	testCase(in0, want0, t)
}
