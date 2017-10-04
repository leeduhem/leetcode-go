package p010

import "testing"

func testCase(s, p string, r0 bool, t *testing.T) {
	r := isMatch(s, p)
	if (r != r0) {
		t.Errorf("Case %v %v: expected %v, got %v", s, p, r0, r)
	}
}

func TestIsMatch(t *testing.T) {
	testCase("aa", "a", false, t)
	testCase("aa", "aa", true, t)
	testCase("aaa", "aa", false, t)
	testCase("aa", "a*", true, t)
	testCase("aa", ".*", true, t)
	testCase("ab", ".*", true, t)
	testCase("aab", "c*a*b", true, t)
	testCase("ab", ".*c", false, t)
	testCase("aaa", "a*a", true, t)
	testCase("a", "ab*", true, t)
	testCase("abbbcd", "ab*bbbcd", true, t)
	testCase("bbbacbaacacaaaba", "b*c*b*.a.*a*.*.*b*", true, t)
	testCase("aaba", "ab*a*c*a", false, t)
	testCase("ab", ".*..", true, t)
	testCase("a", "ab*a", false, t)
	testCase("", "c*", true, t)
	testCase("", "c*c*", true, t)
}
