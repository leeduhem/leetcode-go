package p132

import (
	"testing"
)

func testCase(in string, want int, t *testing.T) {
	got := minCut(in)
	if got != want {
		t.Errorf("Case %v: expected %v, got %v", in, want, got)
	}
}

func TestMinCut(t *testing.T) {
	in0   := "aab"
	want0 := 1
	testCase(in0, want0, t)
}
