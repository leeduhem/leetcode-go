package p003_1

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	testCase := func(in string, want int) {
		got := lengthOfLongestSubstring(in)
		if got != want {
			t.Errorf("Case %v: expected %v, got %v", in, want, got)
		}
	}

	for _, t := range []struct {
		in   string
		want int
	}{
		{"", 0},
		{"c", 1},
		{"bbbbb", 1},
		{"au", 2},
		{"abcabcbb", 3},
		{"pwwkew", 3},
	} {
		testCase(t.in, t.want)
	}
}
