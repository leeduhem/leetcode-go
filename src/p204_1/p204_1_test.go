package p204_1

import (
	"testing"
)

func TestContPrimes(t *testing.T) {
	testCase := func(in, want int) {
		got := countPrimes(in)
		if got != want {
			t.Errorf("Case %v: expected %v, got %v", in, want, got)
		}
	}

	testCase(0, 0)
	testCase(2, 0)
	testCase(3, 1)
	testCase(499979, 41537)
}
