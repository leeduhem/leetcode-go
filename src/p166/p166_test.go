package p166

import (
	"testing"
)

func TestFractionToDecimal(t *testing.T) {
	testCase := func(in0, in1 int, want string) {
		got := fractionToDecimal(in0, in1)
		if got != want {
			t.Errorf("case %v, %v: expected %v, got %v", in0, in1, want, got)
		}
	}

	testCase(1, 2, "0.5")
	testCase(2, 1, "2")
	testCase(2, 3, "0.(6)")
	testCase(1, 6, "0.1(6)")
	testCase(-50, 8, "-6.25")
}
