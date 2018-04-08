package p212

import (
	"testing"
	. "testing/utils"
)

func TestFindWords(t *testing.T) {
	testCase := func(in0 [][]byte, in1, want []string) {
		got := findWords(in0, in1)
		if !IsStrsEqualWithoutOrder(got, want) {
			t.Errorf("Case %v, %v: expected %v, got %v",
				in0, in1, want, got)
		}
	}

	in0 := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	testCase(in0, []string{"oath", "pea", "eat", "rain"}, []string{"eat", "oath"})
}
