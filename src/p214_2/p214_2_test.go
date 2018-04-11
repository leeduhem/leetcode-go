package p214_2

import (
	"testing"
)

func TestShortestPalindrome(t *testing.T) {
	testCase := func(in, want string) {
		got := shortestPalindrome(in)
		if got != want {
			t.Errorf("Case %v: expected %v, got %v", in, want, got)
		}
	}

	testCase("aacecaaa", "aaacecaaa")
	testCase("abcd", "dcbabcd")
}
