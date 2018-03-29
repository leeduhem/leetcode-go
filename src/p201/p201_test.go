package p201

import (
	"testing"
)

func TestRangeBitwiseAnd(t *testing.T) {
	testCase := func(in0, in1, want int) {
		got := rangeBitwiseAnd(in0, in1)
		if got != want {
			t.Errorf("Case %v, %v: expected %v, got %v", in0, in1, want, got)
		}
	}

	testCase(5, 7, 4)
	testCase(4000000, 2147483646, 0)
}
