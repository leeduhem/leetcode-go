package p056

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

func testCase(in0, out0 []Interval, t *testing.T) {
	out1 := merge(in0)
	if ! IsIntervalEqual(out0, out1) {
		t.Errorf("Case %v: expected %v, got %v", in0, out0, out1)
	}
}

func TestMerge(t *testing.T) {
	in0  := []Interval {
		{1,3}, {2,6}, {8,10}, {15,18},
	}
	out0 := []Interval {
		{1,6}, {8,10}, {15,18},
	}
	testCase(in0, out0, t)
}
