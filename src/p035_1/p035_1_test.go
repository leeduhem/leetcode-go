package p035_1

import (
	"sort"
	"testing"
)

func testCase(in0 []int, in1, out0 int, t *testing.T) {
	out1 := searchInsert(in0, in1)
	if out0 != out1 {
		t.Errorf("Case %v %v: expected %v, got %v", in0, in1, out0, out1)
	}
}

func TestSearchInsert(t *testing.T) {
	in0   := []int {1, 3, 5, 6}

	in11  := 5
	out10 := sort.SearchInts(in0, in11)
	testCase(in0, in11, out10, t)

	in21  := 2
	out20 := sort.SearchInts(in0, in21)
	testCase(in0, in21, out20, t)

	in31  := 7
	out30 := sort.SearchInts(in0, in31)
	testCase(in0, in31, out30, t)

	in41  := 0
	out40 := sort.SearchInts(in0, in41)
	testCase(in0, in41, out40, t)

	in50  := []int {1}
	in51  := 2
	out50 := sort.SearchInts(in50, in51)
	testCase(in50, in51, out50, t)

	in60  := []int {1, 3, 5}
	in61  := 6
	out60 := sort.SearchInts(in60, in61)
	testCase(in60, in61, out60, t)

	in70  := []int {1}
	in71  := 1
	out70 := sort.SearchInts(in70, in71)
	testCase(in70, in71, out70, t)
}
