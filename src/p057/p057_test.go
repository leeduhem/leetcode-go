package p057

import "testing"

type Interval struct {
	Start int
	End   int
}

func IsIntervalEqual(in, in1 []Interval) bool {
	if len(in) != len(in1) {
		return false
	}

	for i,v := range in {
		if v.Start != in1[i].Start || v.End != in1[i].End {
			return false
		}
	}

	return true
}

func testCase(in0 []Interval, in1 Interval, out0 []Interval, t *testing.T) {
	out1 := insert(in0, in1)
	if ! IsIntervalEqual(out0, out1) {
		t.Errorf("Case %v %v: expected %v, got %v", in0, in1, out0, out1)
	}
}

func TestInsert(t *testing.T) {
	in00 := []Interval {
		{1,3}, {6,9},
	}
	in01 := Interval {2,5}
	out0 := []Interval {
		{1,5}, {6,9},
	}
	testCase(in00, in01, out0, t)

	in10 := []Interval {
		{1,2}, {3,5}, {6,7}, {8,10}, {12,16},
	}
	in11 := Interval {4,9}
	out1 := []Interval {
		{1,2}, {3,10}, {12,16},
	}
	testCase(in10, in11, out1, t)
}

