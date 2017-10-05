package p018

import "testing"

func testCase(in10 []int, in11 int, out0 [][]int, t *testing.T) {
	out1 := fourSum(in10, in11)
	if len(out0) != len(out1) {
		t.Fatalf("Case %v %v: expected %v, got %v", in10, in11, out0, out1)
	}

	out0m := make(map[string]int, len(out0))
	for i,ns := range out0 {
		key := ints2key(ns)
		out0m[key] = i
	}

	for _,ns := range out1 {
		key := ints2key(ns)
		if _,ok := out0m[key]; !ok {
			t.Fatalf("Case %v %v: expected %v, got %v", in10, in11, out0, out1)
		}
	}
}

func TestFourSum(t *testing.T) {
	in10 := []int {1, 0, -1, 0, -2, 2}
	in11 := 0
	out0 := [][]int {
		{-1,  0, 0, 1},
		{-2, -1, 1, 2},
		{-2,  0, 0, 2},
	}
	testCase(in10, in11, out0, t)
}
