package p115

import (
	"testing"
)

func testCase(in0, in1 string, want int, t *testing.T) {
	got := numDistinct(in0, in1)
	if got != want {
		t.Errorf("Case %v, %v: expected %v, got %v", in0, in1, want, got)
	}
}

func TestNumDistinct(t *testing.T) {
	in00, in01 := "rabbbit", "rabbit"
	want0 := 3
	testCase(in00, in01, want0, t)
}
