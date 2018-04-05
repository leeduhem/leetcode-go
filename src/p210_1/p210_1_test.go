package p210_1

import (
	"testing"
	. "testing/utils"
)

func TestFindOrder(t *testing.T) {
	testCase := func(in0 int, in1, want [][]int) {
		got := findOrder(in0, in1)
		if !IsOneOf(got, want) {
			t.Errorf("Case %v, %v: expected one of %v, got %v",
				in0, in1, want, got)
		}
	}

	testCase(2, [][]int{{1, 0}}, [][]int{{0, 1}})
	testCase(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
		[][]int{{0, 1, 2, 3}, {0, 2, 1, 3}})
}
