package p088

import (
	"testing"
	. "testing/utils"
)

func testCase(in0 []int, in1 int, in2 []int, in3 int, want []int, t *testing.T) {
	oin0 := DeepCopyInts(in0)
	merge(in0, in1, in2, in3)
	if ! IsIntsEqual(in0, want) {
		t.Errorf("Case %v %v, %v %v: expected %v, got %v", oin0, in1, in2, in3, want, in0)
	}
}

func TestMerge(t *testing.T) {
	in00 := make([]int, 6)
	copy(in00, []int {1,3,5})
	in01 := 3
	in02 := []int {2,4,6}
	in03 := 3
	want0 := []int {1,2,3,4,5,6}
	testCase(in00, in01, in02, in03, want0, t)
}
