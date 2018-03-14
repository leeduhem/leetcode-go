package p187

import (
	"testing"
	. "testing/utils"
)

func TestFindRepeatedDnaSequences(t *testing.T) {
	testCase := func(in string, want []string) {
		got := findRepeatedDnaSequences(in)
		if !IsStrsEqual(got, want) {
			t.Errorf("Case %v: expected %v, got %v", in, want, got)
		}
	}

	testCase("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT", []string{"AAAAACCCCC", "CCCCCAAAAA"})
	testCase("AAAAAAAAAAAA", []string{"AAAAAAAAAA"})
}
