package p211

import (
	"testing"
)

func TestWordDirectionary(t *testing.T) {
	testCase0(t)
	testCase1(t)
}

func check(ci int, cnt *int, want, got bool, t *testing.T) {
	if got != want {
		t.Errorf("Case %v [%v]: expected %v, got %v", ci, *cnt, want, got)
	}
	*cnt++
}

func testCase0(t *testing.T) {
	var (
		cnt       int
		got, want bool
	)
	check := func() {
		check(0, &cnt, want, got, t)
	}

	wd := Constructor()
	wd.AddWord("bad")
	wd.AddWord("dad")
	wd.AddWord("mad")
	got, want = wd.Search("pad"), false
	check()
	got, want = wd.Search("bad"), true
	check()
	got, want = wd.Search(".ad"), true
	check()
	got, want = wd.Search("b.."), true
	check()
}

func testCase1(t *testing.T) {
	var (
		cnt       int
		got, want bool
	)
	check := func() {
		check(1, &cnt, want, got, t)
	}

	wd := Constructor()
	wd.AddWord("a")
	wd.AddWord("a")
	got, want = wd.Search("."), true
	check()
	got, want = wd.Search("a"), true
	check()
	got, want = wd.Search("aa"), false
	check()
	got, want = wd.Search("a"), true
	check()
	got, want = wd.Search(".a"), false
	check()
	got, want = wd.Search("a."), false
	check()
}
