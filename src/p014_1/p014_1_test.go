package p014_1

import	"testing"

func testCase(ss []string, r0 string, t *testing.T) {
	r1 := longestCommonPrefix(ss)
	if r0 != r1 {
		t.Errorf("Case %v: expected %v, got %v", ss, r0, r1)
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	ss1 := []string { "a", "ab" }
	r10 := "a"
	testCase(ss1, r10, t)

	ss2 := []string {}
	r20 := ""
	testCase(ss2, r20, t)

	ss3 := []string {""}
	r30 := ""
	testCase(ss3, r30, t)

	ss4 := []string {"a", "b"}
	r40 := ""
	testCase(ss4, r40, t)

	ss5 := []string {"abab","aba",""}
	r50 := ""
	testCase(ss5, r50, t)
}
