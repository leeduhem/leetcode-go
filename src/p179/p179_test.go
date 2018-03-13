package p179

import (
	"testing"
)

func TestLargestNumber(t *testing.T) {
	testCase := func(in []int, want string) {
		got := largestNumber(in)
		if got != want {
			t.Errorf("Case %v: expected %v, got %v", in, want, got)
		}
	}

	testCase([]int{3, 30, 34, 5, 9}, "9534330")
	testCase([]int{121, 12}, "12121")
	testCase([]int{0, 0}, "0")
}
