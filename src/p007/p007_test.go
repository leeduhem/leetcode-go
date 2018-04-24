package p007

import (
	"testing"
)

func TestReverse(t *testing.T) {
	var tests = []struct {
		in, want int
	}{
		{123, 321},
		{-123, -321},
		{1534236469, 0},
	}

	for _, test := range tests {
		got := reverse(test.in)
		if got != test.want {
			t.Errorf("reverse(%v) = %v, want %v",
				test.in, got, test.want)
		}
	}
}
