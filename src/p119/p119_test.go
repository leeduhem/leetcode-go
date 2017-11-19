package p119

import (
	"testing"
	. "testing/utils"
)

func testCase(in int, want []int, t *testing.T) {
	got := getRow(in)
	if ! IsIntsEqual(got, want) {
		t.Errorf("Case %v: expected %v, got %v", in, want, got)
	}
}

func TestGetRow(t *testing.T) {
	in0 := 3
	want0 := []int {1,3,3,1}
	testCase(in0, want0, t)

	in1 := 21
	want1 := []int {
		1,21,210,1330,5985,20349,54264,116280,203490,293930,352716,352716,293930,203490,116280,54264,20349,5985,1330,210,21,1,
	}
	testCase(in1, want1, t)
}
