package p071

import (
	"testing"
)

func testCase(in0, out0 string, t *testing.T) {
	out1 := simplifyPath(in0)
	if out0 != out1 {
		t.Errorf("Case %v: expected %v, got %v", in0, out0, out1)
	}
}

func TestSimplifyPath(t *testing.T) {
	cases := []struct {
		in, want string
	} {
		{"/home/", "/home"},
		{"/a/./b/../../c/", "/c"},
		{"/../", "/"},
		{"/home//foo/", "/home/foo"},
	}
	for _,c := range cases {
		testCase(c.in, c.want, t)
	}
}
