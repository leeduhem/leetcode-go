package p040

import (
	"sort"
	"strconv"
	"strings"
	"testing"
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
	out1 := combinationSum2(in0, in1)
	if len(out0) != len(out1) {
		t.Fail()
	}

	if !t.Failed() {
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

func TestCombinationSum2(t *testing.T) {
	in10  := []int {10, 1, 2, 7, 6, 1, 5}
	in11  := 8
	out10 := [][]int {
		{1, 7},
		{1, 2, 5},
		{2, 6},
		{1, 1, 6},
	}
	testCase(in10, in11, out10, t)
}
