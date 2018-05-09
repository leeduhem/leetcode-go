package p009

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		in   int
		want bool
	}{
		{1, true},
		{-1, false},
	}

	for _, test := range tests {
		got := isPalindrome(test.in)
		if got != test.want {
			t.Errorf("isPalindrome(%v) = %v, want = %v",
				test.in, got, test.want)
		}
	}
}
