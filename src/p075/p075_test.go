package p075

import (
	"testing"
	. "testing/utils"
)

const (
	red = iota
	white
	blue
)

func testCase(in0, want []int, t *testing.T) {
	got := DeepCopyInts(in0)
	sortColors(got)
	if ! IsIntsEqual(want, got) {
		t.Errorf("Case %v: expected %v, got %v", in0, want, got)
	}
}

func TestSortColors(t *testing.T) {
	in0   := []int {white, red, blue, blue, red, white}
	want0 := []int {red, red, white, white, blue, blue}
	testCase(in0, want0, t)
}
