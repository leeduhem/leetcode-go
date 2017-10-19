package p045

import "testing"

func testCase(in0 []int, out0 int, t *testing.T) {
	out1 := jump(in0)
	if out0 != out1 {
		t.Errorf("Case %v: expected %v, got %v", in0, out0, out1)
	}
}

func TestJump(t *testing.T) {
	in1  := []int {2,3,1,1,4}
	out1 := 2
	testCase(in1, out1, t)

	in2  := []int {0}
	out2 := 0
	testCase(in2, out2, t)

	in3  := []int {1}
	out3 := 0
	testCase(in3, out3, t)

	in4  := []int {1,1,1,1}
	out4 := 3
	testCase(in4, out4, t)

	in5  := []int {1,2,1,1,1}
	out5 := 3
	testCase(in5, out5, t)

	in6  := []int {10,9,8,7,6,5,4,3,2,1,1,0}
	out6 := 2
	testCase(in6, out6, t)

	in7  := []int {1,3,2}
	out7 := 2
	testCase(in7, out7, t)

	in8  := []int {1,2,3}
	out8 := 2
	testCase(in8, out8, t)

	in9  := []int {1,2,0,1}
	out9 := 2
	testCase(in9, out9, t)

	in10  := []int {4,1,1,3,1,1,1}
	out10 := 2
	testCase(in10, out10, t)

	in11  := []int {7,0,9,6,9,6,1,7,9,0,1,2,9,0,3}
	out11 := 2
	testCase(in11, out11, t)
}
