package p080

import (
	"testing"
)

func testCase(in0 []int, want int, t *testing.T) {
	got := removeDuplicates(in0)
	if got != want {
		t.Errorf("Case %v: expected %v, got %v", in0, want, got)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	in0 := []int {1,1,1,2,2,3}
	want0 := 5
	testCase(in0, want0, t)
}
