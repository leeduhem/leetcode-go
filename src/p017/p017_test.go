package p017

import "testing"

func testCase(in1 string, out0 []string, t *testing.T) {
	out1 := letterCombinations(in1)
	if len(out0) != len(out1) {
		t.Fatalf("Case %v: expected %v, got %v", in1, out0, out1)
	}

	out0m := make(map[string]int)
	for i,v := range out0 {
		out0m[v] = i
	}

	for _,v := range out1 {
		if _,ok := out0m[v]; !ok {
			t.Fatalf("Case %v: expected %v, got %v", in1, out0, out1)
		}
	}
}

func TestLetterCombinations(t *testing.T) {
	in1  := "23"
	out0 := []string {"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	testCase(in1, out0, t)
}
