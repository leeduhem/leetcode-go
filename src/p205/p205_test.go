package p205

import (
	"testing"
)

func TestIsIsomorphic(t *testing.T) {
	testCase := func(in0, in1 string, want bool) {
		got := isIsomorphic(in0, in1)
		if got != want {
			t.Errorf("Case %v, %v: expected %v, got %v", in0, in1, want, got)
		}
	}

	testCase("egg", "add", true)
	testCase("foo", "bar", false)
	testCase("paper", "title", true)

	testCase("ab", "aa", false)
}
