package p140

import (
	"testing"
	. "testing/utils"
)

func testCase(in0 string, in1, want []string, t *testing.T) {
	got := wordBreak(in0, in1)
	if ! IsStrsEqualWithoutOrder(want, got) {
		t.Errorf("Case %v %#v: expected %#v, got %#v", in0, in1, want, got)
	}
}

func TestWordBreak(t *testing.T) {
	in00 := "catsanddog"
	in01 := []string { "cat", "cats", "and", "sand", "dog" }
	want0 := []string { "cats and dog", "cat sand dog" }
	testCase(in00, in01, want0, t)
}
