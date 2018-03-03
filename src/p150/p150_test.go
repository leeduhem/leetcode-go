package p150

import (
	"testing"
)

func testCase(in []string, want int, t *testing.T) {
	got := evalRPN(in)
	if got != want {
		t.Errorf("Case %v: expected %v, got %v", in, want, got)
	}
}

func TestEvalRPN(t *testing.T) {
	in0 := []string { "2", "1", "+", "3", "*" }
	want0 := 9
	testCase(in0, want0, t)

	in1 := []string { "4", "13", "5", "/", "+" }
	want1 := 6
	testCase(in1, want1, t)
}
