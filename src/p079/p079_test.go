package p079

import (
	"testing"
	"fmt"
)

type board [][]byte

func (b board) String() string {
	s := "\n["
	for i, r := range b {
		if i > 0 {
			s += " "
		}

		s += "["
		s += fmt.Sprintf("%c", r[0])
		for j := 1; j < len(r); j++ {
			s += fmt.Sprintf(" %c", r[j])
		}
		s += "]"

		if i < len(b) -1 {
			s += "\n"
		}
	}
	s += "]"

	return s
}

func testCase(in0 [][]byte, in1 string, want bool, t *testing.T) {
	got := exist(in0, in1)
	if got != want {
		t.Errorf("Case %v, %v: expected %v, got %v", board(in0), in1, want, got)
	}
}

func TestExist(t *testing.T) {
	in00 := [][]byte {
		{'A','B','C','E'},
		{'S','F','C','S'},
		{'A','D','E','E'},
	}
	in01 := "ABCCED"
	want0 := true
	testCase(in00, in01, want0, t)

	in11 := "SEE"
	want1 := true
	testCase(in00, in11, want1, t)

	in21 := "ABCB"
	want2 := false
	testCase(in00, in21, want2, t)
}
