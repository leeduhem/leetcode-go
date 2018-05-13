package p010_1

import (
	"testing"
)

func TestIsMatch(t *testing.T) {
	tests := []struct {
		in0, in1 string
		want     bool
	}{
		{"aa", "a", false},
		{"aa", "aa", true},
		{"aaa", "aa", false},
		{"aa", "a*", true},
		{"aa", ".*", true},
		{"ab", ".*", true},
		{"aab", "c*a*b", true},
		{"ab", ".*c", false},
		{"aaa", "a*a", true},
		{"a", "ab*", true},
		{"abbbcd", "ab*bbbcd", true},
		{"bbbacbaacacaaaba", "b*c*b*.a.*a*.*.*b*", true},
		{"aaba", "ab*a*c*a", false},
		{"ab", ".*..", true},
		{"a", "ab*a", false},
		{"", "c*", true},
		{"", "c*c*", true},
	}

	for _, test := range tests {
		got := isMatch(test.in0, test.in1)
		if got != test.want {
			t.Errorf("isMatch(%q, %q) = %v, want %v",
				test.in0, test.in1, got, test.want)
		}
	}
}
