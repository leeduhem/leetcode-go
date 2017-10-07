package p026

import "testing"

func testCase(in0 []int, out0 int, t *testing.T) {
	out1 := removeDuplicates(in0)
	if out0 != out1 {
		t.Fail()
	}

	out1m := make(map[int]int, out1)
	for i := 0; i < out1; i++ {
		v := in0[i]
		if _,ok := out1m[v]; ok {
			t.Fail()
		} else {
			out1m[v] = i
		}
	}

	if t.Failed() {
		t.Errorf("Case %v: expected %v, got %v", in0, out0, out1)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	in0  := []int {1, 1, 2}
	out0 := 2
	testCase(in0, out0, t)

	in1  := []int {1, 1}
	out1 := 1
	testCase(in1, out1, t)
}
