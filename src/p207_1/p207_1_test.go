package p207_1

import (
	"testing"
)

func TestCanFinish(t *testing.T) {
	testCase := func(in0 int, in1 [][]int, want bool) {
		got := canFinish(in0, in1)
		if got != want {
			t.Errorf("Case %v, %v: expected %v, got %v", in0, in1, want, got)
		}
	}

	testCase(2, [][]int{{1, 0}}, true)
	testCase(2, [][]int{{1, 0}, {0, 1}}, false)
	testCase(3, [][]int{{1, 0}, {0, 2}, {2, 1}}, false)
	testCase(3, [][]int{{1, 0}, {2, 0}, {0, 2}}, false)
}
