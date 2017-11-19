package p118

import (
	"testing"
	. "testing/utils"
)

func testCase(in int, want [][]int, t *testing.T) {
	got := generate(in)
	if ! Is2DIntsEqual(got, want) {
		t.Errorf("Case %v: expected %v, got %v", in, want, got)
	}
}

func TestGenerate(t *testing.T) {
	in0 := 5
	want0 := [][]int {
		{1},
		{1,1},
		{1,2,1},
		{1,3,3,1},
		{1,4,6,4,1},
	}
	testCase(in0, want0, t)
}
