package p049

import "testing"

func IsStrSliceEqual(ss, ss1 []string) bool {
	m := map[string]int {}
	for i,s := range ss {
		m[s] = i
	}

	for _,s := range ss1 {
		delete(m, s)
	}

	return len(m) == 0
}

func testCase(in0 []string, out0 [][]string, t *testing.T) {
	out1 := groupAnagrams(in0)
	if len(out0) != len(out1) {
		t.Fail()
	}

	if ! t.Failed() {
	OuterLoop:
		for _,ss := range out1 {
			for _,ss1 := range out0 {
				if IsStrSliceEqual(ss, ss1) {
					continue OuterLoop
				}
			}
			t.Fail()
			break
		}
	}

	if t.Failed() {
		t.Errorf("Case %v: expected %v, got %v", in0, out0, out1)
	}
}

func TestGroupAnagrams(t *testing.T) {
	in1  := []string {"eat", "tea", "tan", "ate", "nat", "bat"}
	out1 := [][]string {
		{"ate", "eat","tea"},
		{"nat","tan"},
		{"bat"},
	}
	testCase(in1, out1, t)
}

