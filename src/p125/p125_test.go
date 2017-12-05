package p125

import (
	"testing"
)

func testCase(in string, want bool, t *testing.T) {
	got := isPalindrome(in)
	if got != want {
		t.Errorf("Case %v: expected %v, got %v", in, want, got)
	}
}

func TestIsPalindrome(t *testing.T) {
	in0 := "A man, a plan, a canal: Panama"
	want0 := true
	testCase(in0, want0, t)

	in1 := "race a car"
	want1 := false
	testCase(in1, want1, t)

	in2 := ".,"
	want2 := true
	testCase(in2, want2, t)
}
