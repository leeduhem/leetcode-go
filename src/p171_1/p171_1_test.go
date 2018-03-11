package p171_1

import (
	"testing"
)

func TestTitleToNumber(t *testing.T) {
	testCase := func(in string, want int) {
		got := titleToNumber(in)
		if got != want {
			t.Errorf("Case %v: expected %v, got %v", in, want, got)
		}
	}

	testCase("A", 1)
	testCase("B", 2)
	testCase("Z", 26)
	testCase("AA", 27)
	testCase("AB", 28)
}
