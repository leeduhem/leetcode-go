package p008

import "testing"

func testCase(r0, r1 int, t *testing.T) {
	if (r0 != r1) {
		t.Errorf("Expected %v, got %v", r0, r1)
	}
}

func TestMyAtoi(t *testing.T) {
	testCase(10, myAtoi("10"), t)
	testCase(123, myAtoi(" 123"), t)
	testCase(-123, myAtoi("  -123"), t)
	testCase(0, myAtoi(""), t)
	testCase(0, myAtoi("+"), t)
	testCase(2147483647, myAtoi("2147483648"), t)
	testCase(2147483647, myAtoi("9223372036854775809"), t)

}
