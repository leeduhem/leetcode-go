package p154

import (
	"testing"
)

func testCase(in []int, want int, t *testing.T) {
	got := findMin(in)
	if got != want {
		t.Errorf("Case %v: expected %v, got %v", in, want, got)
	}
}

func TestFindMin(t *testing.T) {
	in0 := []int {4,5,6,7,0,1,2}
	want0 := 0
	testCase(in0, want0, t)

	in1 := []int { 10,1,10,10,10 }
	want1 := 1
	testCase(in1, want1, t)
}
