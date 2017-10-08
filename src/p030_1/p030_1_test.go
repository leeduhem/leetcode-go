package p030_1

import "testing"

func testCase(in0 string, in1 []string, out0 []int, t *testing.T) {
	out1 := findSubstring(in0, in1)
	if len(out0) != len(out1) {
		t.Fail()
	}

	out0m := make(map[int]int, len(out0))
	for i,v := range out0 {
		out0m[v] = i
	}

	for _,v := range out1 {
		if _,ok := out0m[v]; !ok {
			t.Fail()
		}
	}

	if t.Failed() {
		t.Errorf("Case %v %v: expected %v, got %v", in0, in1, out0, out1)
	}
}

func TestFindSubstring(t *testing.T) {
	in10  := "barfoothefoobarman"
	in11  := []string {"foo", "bar"}
	out10 := []int {0, 9}
	testCase(in10, in11, out10, t)
}
