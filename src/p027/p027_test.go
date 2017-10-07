package p027

import "testing"

func testCase(in0 []int, in1, out0 int, t *testing.T) {
	out1 := removeElement(in0, in1)
	if out0 != out1 {
		t.Fail()
	}

	for i := 0; i < out1; i++ {
		if in0[i] == in1 {
			t.Fail()
		}
	}

	if t.Failed() {
		t.Errorf("Case %v %v: expected %v, got %v", in0, in1, out0, out1)
	}
}

func TestRemoveElement(t *testing.T) {
	in10  := []int {3, 2, 2, 3}
	in11  := 3
	out10 := 2
	testCase(in10, in11, out10, t)

	in20  := []int {1}
	in21  := 1
	out20 := 0
	testCase(in20, in21, out20, t)

	in30  := []int {3, 3}
	in31  := 5
	out30 := 2
	testCase(in30, in31, out30, t)

	in40  := []int {2}
	in41  := 3
	out40 := 1
	testCase(in40, in41, out40, t)
}
