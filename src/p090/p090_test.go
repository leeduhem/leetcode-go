package p090

import (
	"testing"
	. "testing/utils"
)

func testCase(in0 []int, want [][]int, t *testing.T) {
	got := subsetsWithDup(in0)
	if ! Is2DIntsEqualWithoutOrder(want, got) {
		t.Errorf("Case %v: expected %v, got %v", in0, want, got)
	}
}

func TestSubsets(t *testing.T) {
	in0   := []int {1,2,2}
	want0 := [][]int {
		{2},
		{1},
		{1,2,2},
		{2,2},
		{1,2},
		{},
	}
	testCase(in0, want0, t)
}
