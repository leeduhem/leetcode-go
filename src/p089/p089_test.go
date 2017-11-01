package p089

import (
	"testing"
	. "testing/utils"
)

func testCase(in int, want []int, t *testing.T) {
	got := grayCode(in)
	if ! IsIntsEqual(got, want) {
		t.Errorf("Case %v: expected %v, got %v", in, want, got)
	}
}

func TestGrayCode(t *testing.T) {
	in0 := 2
	want0 := []int {0,1,3,2}
	testCase(in0, want0, t)

	in1 := 3
	want1 := []int {0,1,3,2,6,7,5,4}
	testCase(in1, want1, t)
}
