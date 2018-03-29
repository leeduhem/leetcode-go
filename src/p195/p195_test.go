package p195

import (
	"testing"
)

func TestTenthLine(t *testing.T) {
	testCase := func(in, want string) {
		got := tenthLine(in)
		if got != want {
			t.Errorf("Case %v: expected %v, got %v", in, want, got)
		}
	}

	testCase("file.txt", "Line 10")
	testCase("file1.txt", "")
}
