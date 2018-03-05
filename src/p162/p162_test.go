package p162

import (
	"testing"
)

func isIn(e int, lst []int) bool {
	for _, l := range lst {
		if l == e {
			return true
		}
	}

	return false
}

func testCase(in, want []int, t *testing.T) {
	got := findPeakElement(in)
	if ! isIn(got, want) {
		t.Errorf("Case %v: expected one of %v, got %v", in, want, got)
	}
}

func TestFindPeakElement(t *testing.T) {
	in0 := []int {1,2,3,1}
	want0 := []int {2}
	testCase(in0, want0, t)
}
