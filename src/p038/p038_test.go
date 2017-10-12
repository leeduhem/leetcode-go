package p038

import "testing"

func testCase(in0 int, out0 string, t *testing.T) {
	out1 := countAndSay(in0)
	if out0 != out1 {
		t.Errorf("Case %v: expected %v, got %v", in0, out0, out1)
	}
}

func TestCountAndSay(t *testing.T) {
	in10  := 1
	out10 := "1"
	testCase(in10, out10, t)

	in20  := 4
	out20 := "1211"
	testCase(in20, out20, t)
}
