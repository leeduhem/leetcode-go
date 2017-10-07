package p022

import "testing"

func testCase(in0 int, out0 []string, t *testing.T) {
	out1 := generateParenthesis(in0)

	if len(out0) != len(out1) {
		t.Fatalf("Case %v: expected %v, got %v", in0, out0, out1)
	}

	out0m := make(map[string]int, len(out0))
	for i,s := range out0 {
		out0m[s] = i
	}

	for _,s := range out1 {
		if _,ok := out0m[s]; !ok {
			t.Fatalf("Case %v: expected %v, got %v", in0, out0, out1)
		}
	}
}

func TestGenerateParenthesis(t *testing.T) {
	in10  := 3
	out10 := []string {
		"((()))",
		"(()())",
		"(())()",
		"()(())",
		"()()()",
	}
	testCase(in10, out10, t)
}
