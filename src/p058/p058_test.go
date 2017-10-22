package p058

import "testing"

func testCase(in0 string, out0 int, t *testing.T) {
	out1 := lengthOfLastWord(in0)
	if out0 != out1 {
		t.Errorf("Case %v: expected %v, got %v", in0, out0, out1)
	}
}

func TestLengthOfLastWord(t *testing.T) {
	in0  := "Hello World"
	out0 := 5
	testCase(in0, out0, t)

	in1  := "a "
	out1 := 1
	testCase(in1, out1, t)
}
