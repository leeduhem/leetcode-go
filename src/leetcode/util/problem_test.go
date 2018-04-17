package util

import (
	"testing"
)

func TestPid2Pkg(t *testing.T) {
	testCase := func(in int, want string) {
		got := Pid2Pkg(in)
		if got != want {
			t.Errorf("Case TestPid2Pkg(%d): expected %v, got %v",
				in, want, got)
		}
	}

	testCase(1, "p001")
	testCase(12, "p012")
	testCase(123, "p123")
	testCase(1234, "p1234")
}
