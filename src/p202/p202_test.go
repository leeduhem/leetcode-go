package p202

import (
	"testing"
)

func TestIsHappy(t *testing.T) {
	testCase := func(in int, want bool) {
		got := isHappy(in)
		if got != want {
			t.Errorf("Case %v: expected %v, got %v", in, want, got)
		}
	}

	testCase(19, true)
}
