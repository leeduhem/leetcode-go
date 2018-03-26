package p192

import (
	"testing"
	. "testing/utils"
)

func TestWordFrequency(t *testing.T) {
	testCase := func(in string, want []string) {
		got := wordFrequency(in)
		if !IsStrsEqual(got, want) {
			t.Errorf("Case %v: expected %v, got %v", in, want, got)
		}
	}

	testCase("words.txt", []string{
		"the 4",
		"is 3",
		"sunny 2",
		"day 1",
	})
}
