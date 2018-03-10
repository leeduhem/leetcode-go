package p168

import (
	"testing"
)

func TestConvertToTitle(t *testing.T) {
	testCase := func(in int, want string) {
		got := convertToTitle(in)
		if got != want {
			t.Errorf("Case %v: expected %v, got %v", in, want, got)
		}
	}

	testCase(1, "A")
	testCase(2, "B")
	testCase(26, "Z")
	testCase(27, "AA")
	testCase(28, "AB")

	testCase(52, "AZ")
}
