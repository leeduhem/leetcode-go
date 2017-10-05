package p003

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	ss := []string { "", "c", "bbbbb", "au", "abcabcbb", "pwwkew" }
	rs := []int    {  0,  1,   1,       2,    3,          3       }
	for i,s := range ss {
		max := lengthOfLongestSubstring(s)
		if max != rs[i] {
			t.Errorf("Case %v: expected %v, got %v", s, rs[i], max)
		}
	}
}

