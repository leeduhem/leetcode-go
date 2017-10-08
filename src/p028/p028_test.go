package p028

import (
	"testing"
	"strings"
)

func testCase(in0, in1 string, out0 int, t *testing.T) {
	out1 := strStr(in0, in1)
	if out0 != out1 {
		t.Fail()
	}

	if out0 != -1 {
		for i := 0; i < len(in1); i++ {
			if out0+i < len(in0) && in0[out0+i] != in1[i] {
				t.Fail()
			}
		}
	}

	if t.Failed() {
		t.Errorf("Case %v %v: expected %v, got %v", in0, in1, out0, out1)
	}
}

func TestStrStr(t *testing.T) {
	in10 := "abc"
	in11 := "b"
	out10 := strings.Index(in10, in11)
	testCase(in10, in11, out10, t)

	in20 := "abc"
	in21 := "abc"
	out20 := strings.Index(in20, in21)
	testCase(in20, in21, out20, t)

	in30 := "中ab英cd混ef杂"
	in31 := "混"
	out30 := strings.Index(in30, in31)
	testCase(in30, in31, out30, t)

	in40 := ""
	out40 := strings.Index(in40, in40)
	testCase(in40, in40, out40, t)

	in50 := "a"
	in51 := ""
	out50 := strings.Index(in50, in51)
	testCase(in50, in51, out50, t)

	in60 := "mississippi"
	in61 := "issipi"
	out60 := strings.Index(in60, in61)
	testCase(in60, in61, out60, t)
}
