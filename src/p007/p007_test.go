package p007

import "testing"

func testCase (n, r0 int, t *testing.T) {
	r1 := reverse(n)
	if (r0 != r1) {
		t.Errorf("Expected %v, got %v", r0, r1)
	}
}

func TestReverse(t *testing.T) {
	testCase(123, 321, t)
	testCase(-123, -321, t)
	testCase(1534236469, 0, t)
}
