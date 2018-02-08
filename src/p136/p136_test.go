package p136

import (
	"testing"
)

func testCase(in []int, want int, t *testing.T) {
	got := singleNumber(in)
	if got != want {
		t.Errorf("Case %v: expected %v, got %v", in, want, got)
	}
}

func TestSingleNumber(t *testing.T) {
	in0 := []int {1}
	want0 := 1
	testCase(in0, want0, t)

	in1 := []int {2,2,1}
	want1 := 1
	testCase(in1, want1, t)

	in2 := []int {1,0,1}
	want2 := 0
	testCase(in2, want2, t)
}
