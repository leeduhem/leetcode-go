package p034

import "testing"

func testCase(in0 []int, in1 int, out0 []int, t *testing.T) {
	out1 := searchRange(in0, in1)
	if len(out0) != len(out1) {
		t.Fail()
	}

	if ! t.Failed() {
		for i,v := range out0 {
			if v != out1[i] {
				t.Fail()
			}
		}
	}

	if t.Failed() {
		t.Errorf("Case %v %v: expected %v, got %v", in0, in1, out0, out1)
	}
}

func TestSearchRange(t *testing.T) {
	in10  := []int {5, 7, 7, 8, 8, 10}
	in11  := 8
	out10 := []int {3, 4}
	testCase(in10, in11, out10, t)

	in20  := []int {1}
	in21  := 1
	out20 := []int {0, 0}
	testCase(in20, in21, out20, t)

	in30  := []int {}
	in31  := 0
	out30 := []int {-1, -1}
	testCase(in30, in31, out30, t)

	in40  := []int {1}
	in41  := 0
	out40 := []int {-1, -1}
	testCase(in40, in41, out40, t)
}
