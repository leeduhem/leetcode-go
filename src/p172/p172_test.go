package p172

import (
	"testing"
)

func TestTrailingZeroes(t *testing.T) {
	testCase := func(in, want int) {
		got := trailingZeroes(in)
		if got != want {
			t.Errorf("Case %v: expected %v, got %v", in, want, got)
		}
	}

	testCase(3, 0)
	testCase(5, 1)
	testCase(10, 2)
	testCase(80, 19)
	testCase(625, 156)
}
