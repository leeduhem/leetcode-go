package p068

import (
	"fmt"
	"testing"
)
import . "testing/utils"

type paragraph []string

func (p paragraph) String() string {
	str := "["
	for i,s := range p {
		if i > 0 {
			str += " "
		}

		str += fmt.Sprintf("[%v]", s)

		if i < len(p)-1 {
			str += "\n"
		}
	}
	str += "]"

	return str
}

func testCase(in0 []string, in1 int, out0 []string, t *testing.T) {
	out1 := fullJustify(in0, in1)
	if ! IsStrsEqual(out0, out1) {
		t.Errorf("Case %v %v: expected\n%v, got\n%v", in0, in1, paragraph(out0), paragraph(out1))
	}
}

func TestFullJustify(t *testing.T) {
	in00 := []string {"This", "is", "an", "example", "of", "text", "justification."}
	in01 := 16
	out0 := []string {
		"This    is    an",
		"example  of text",
		"justification.  ",
	}
	testCase(in00, in01, out0, t)

	in10 := []string {"Listen","to","many,","speak","to","a","few."}
	in11 := 6
	out1 := []string {
		"Listen",
		"to    ",
		"many, ",
		"speak ",
		"to   a",
		"few.  ",
	}
	testCase(in10, in11, out1, t)
}
