package p128

import (
	"testing"
)

func testCase(in []int, want int, t *testing.T) {
	got := longestConsecutive(in)
	if got != want {
		t.Errorf("Case %v: expected %v, got %v", in, want, got)
	}
}

func TestLongestConsecutive(t *testing.T) {
	in0 := []int { 100, 4, 200, 1, 3, 2 }
	want0 := 4
	testCase(in0, want0, t)
}
