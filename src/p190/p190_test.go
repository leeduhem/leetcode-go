package p190

import (
	"testing"
)

func TestReverseBits(t *testing.T) {
	testCase := func(in, want uint32) {
		got := reverseBits(in)
		if got != want {
			t.Errorf("Case %v: expected %v, got %v", in, want, got)
		}
	}

	testCase(43261596, 964176192)
}
