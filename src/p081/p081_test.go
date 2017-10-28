package p081

import (
	"testing"
)

func testCase(in0 []int, in1 int, want bool, t *testing.T) {
	got := search(in0, in1)
	if got != want {
		t.Errorf("Case %v, %v: expected %v, got %v", in0, in1, want, got)
	}
}

func TestSearch(t *testing.T) {
	in00 := []int {4,5,6,7,0,1,2}
	in01 := 7
	want0 := true
	testCase(in00, in01, want0, t)
}
