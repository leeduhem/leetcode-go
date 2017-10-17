package p039_1

import (
	"testing"
	"sort"
	"strconv"
	"strings"
)

func ints2key(ns []int) string {
	sort.Ints(ns)

	ss := []string {}
	for _,n := range ns {
		ss = append(ss, strconv.Itoa(n))
	}

	return strings.Join(ss, ",")
}

func testCase(in0 []int, in1 int, out0 [][]int, t *testing.T) {
	out1 := combinationSum(in0, in1)
	if len(out1) != len(out0) {
		t.Fail()
	}

	if ! t.Failed() {
		out0m := map[string]int {}
		for i,v := range out0 {
			key := ints2key(v)
			out0m[key] = i
		}

		for _,v := range out1 {
			key := ints2key(v)
			if _,ok := out0m[key]; !ok {
				t.Fail()
				break
			}
		}
	}

	if t.Failed() {
		t.Errorf("Case %v, %v: expected %v, got %v", in0, in1, out0, out1)
	}
}

func TestCombinationSum(t *testing.T) {
	in10  := []int {2, 3, 6, 7}
	in11  := 7
	out10 := [][]int {
		{7},
		{2,2,3},
	}
	testCase(in10, in11, out10, t)

	in20  := []int {1, 2}
	in21  := 2
	out20 := [][]int {
		{2},
		{1,1},
	}
	testCase(in20, in21, out20, t)

	in30  := []int {7, 3, 2}
	in31  := 18
	out30 := [][]int {
		{2, 2, 2, 2, 2, 2, 2, 2, 2},
		{2, 2, 2, 2, 2, 2, 3, 3},
		{2, 2, 2, 2, 3, 7},
		{2, 2, 2, 3, 3, 3, 3},
		{2, 2, 7, 7},
		{2, 3, 3, 3, 7},
		{3, 3, 3, 3, 3, 3},
	}
	testCase(in30, in31, out30, t)
}
