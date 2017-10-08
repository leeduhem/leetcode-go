package p029_1

import "testing"

func testCase(in0, in1, out0 int, t *testing.T) {
	out1 := divide(in0, in1)
	if out0 != out1 {
		t.Errorf("Case %v %v: expected %v, got %v", in0, in1, out0, out1)
	}
}

func TestDivide(t *testing.T) {
	in10  := 10
	in11  := 3
	out10 := in10 / in11
	testCase(in10, in11, out10, t)

	in20  := 1
	out20 := in20 / in20
	testCase(in20, in20, out20, t)

	in30  := -2147483648
	in31  := 1
	out30 := in30 / in31
	testCase(in30, in31, out30, t)

	in40  := -2147483648
	in41  := -1
	// Looks like leetcode is supposed that int is int32
	out40 := 2147483647
	testCase(in40, in41, out40, t)

	in50  := -2147483648
	in51  := 2
	out50 := in50 / in51
	testCase(in50, in51, out50, t)
}
