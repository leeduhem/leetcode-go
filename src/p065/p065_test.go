package p065

import (
	"testing"
)

func testCase(in0 string, out0 bool, t *testing.T) {
	out1 := isNumber(in0)
	if out0 != out1 {
		t.Errorf("Case %v: expected %v, got %v", in0, out0, out1)
	}
}

func TestIsNumber(t *testing.T) {
	in0s  := []string {"0", " 0.1 ", "abc", "1 a", "2e10"}
	out0s := []bool   {true,   true, false, false,   true}
	for i, in0 := range in0s {
		testCase(in0, out0s[i], t)
	}

	in1  := "."
	out1 := false
	testCase(in1, out1, t)

	in2  := "0e"
	out2 := false
	testCase(in2, out2, t)

	in3  := "3. "
	out3 := true
	testCase(in3, out3, t)

	in4  := "0.."
	out4 := false
	testCase(in4, out4, t)

	in5  := ".1."
	out5 := false
	testCase(in5, out5, t)

	in6  := "6+1"
	out6 := false
	testCase(in6, out6, t)

	in7  := "6e6.5"
	out7 := false
	testCase(in7, out7, t)

	in8  := "92e1740e91"
	out8 := false
	testCase(in8, out8, t)

	in9  := "459277e38+"
	out9 := false
	testCase(in9, out9, t)
}
