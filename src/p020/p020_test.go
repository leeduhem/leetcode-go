package p020

import "testing"

func testCase(in0 string, out0 bool, t *testing.T) {
	out1 := isValid(in0)
	if out0 != out1 {
		t.Errorf("Case %v: expected %v, got %v", in0, out0, out1)
	}
}

func TestIsValid(t *testing.T) {
	in1 := "()"
	testCase(in1, true, t)

	in2 := "()[]{}"
	testCase(in2, true, t)

	in3 := "(]"
	testCase(in3, false, t)

	in4 := "([)]"
	testCase(in4, false, t)
}
