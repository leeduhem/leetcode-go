package p191

import (
	"testing"
)

func TestHammingWeight(t *testing.T) {
	testCase := func(in uint32, want int) {
		got := hammingWeight(in)
		if got != want {
			t.Errorf("Case %v: expected %v, got %v", in, want, got)
		}
	}

	testCase(11, 3)
	testCase(2, 1)
}
