package p078

import (
	"testing"
	. "testing/utils"
)

func testCase(in0 []int, want [][]int, t *testing.T) {
	got := subsets(in0)
	if ! Is2DIntsEqualWithoutOrder(want, got) {
		t.Errorf("Case %v: expected %v, got %v", in0, want, got)
	}
}

func TestSubsets(t *testing.T) {
	in0   := []int {1,2,3}
	want0 := [][]int {
		{3},
		{1},
		{2},
		{1,2,3},
		{1,3},
		{2,3},
		{1,2},
		{},
	}
	testCase(in0, want0, t)
}
