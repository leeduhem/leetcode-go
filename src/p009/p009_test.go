package p009

import "testing"

func testCase(n int, r bool, t *testing.T) {
	r1 := isPalindrome(n)
	if (r != r1) {
		t.Errorf("Case %v: expect %v, got %v", n, r, r1)
	}
}

func TestIsPalindrome(t *testing.T) {
	testCase(1, true, t)
	testCase(-1, false, t)
}
