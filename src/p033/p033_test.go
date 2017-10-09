package p033

import "testing"

func testCase(in0 []int, in1, out0 int, t *testing.T) {
	out1 := search(in0, in1)
	if out0 != out1 || (out0 >= 0 && in0[out0] != in1) {
		t.Fatalf("Case %v %v: expected %v, got %v", in0, in1, out0, out1)
	}
}

func TestSearch(t *testing.T) {
	in10  := []int {4, 5, 6, 7, 0, 1, 2}
	in11  := 6
	out10 := 2
	testCase(in10, in11, out10, t)

	in20  := []int {1}
	in21  := 0
	out20 := -1
	testCase(in20, in21, out20, t)

	in30  := []int {1, 3}
	in31  := 2
	out30 := -1
	testCase(in30, in31, out30, t)

	in40  := []int {3, 1}
	in41  := 1
	out40 := 1
	testCase(in40, in41, out40, t)

	in50  := []int {1, 3}
	in51  := 1
	out50 := 0
	testCase(in50, in51, out50, t)

	in60  := []int {1}
	in61  := 2
	out60 := -1
	testCase(in60, in61, out60, t)
}
