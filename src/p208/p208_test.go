package p208

import (
	"testing"
)

func check(ci int, cnt *int, want, got bool, t *testing.T) {
	if got != want {
		t.Errorf("Case %v [%v]: expected %v, got %v", ci, *cnt, want, got)
	}
	*cnt++
}

func testCase0(t *testing.T) {
	var (
		cnt int
		got, want bool
	)
	check := func() {
		check(0, &cnt, want, got, t)
	}

	trie := Constructor()
	got, want = trie.Search("a"), false; check()
}

func testCase1(t *testing.T) {
	var (
		cnt int
		got, want bool
	)
	check := func() {
		check(1, &cnt, want, got, t)
	}

	trie := Constructor()
	trie.Insert("ab")
	got, want = trie.Search("a"), false; check()
	got, want = trie.StartsWith("a"), true; check()
}

func testCase2(t *testing.T) {
	var (
		cnt int
		got, want bool
	)
	check := func() {
		check(2, &cnt, want, got, t)
	}

	trie := Constructor()
	trie.Insert("abc")
	got, want = trie.Search("abc"), true; check()
	got, want = trie.Search("ab"), false; check()
	trie.Insert("ab")
	got, want = trie.Search("ab"), true; check()
	trie.Insert("ab")
	got, want = trie.Search("ab"), true; check()
}

func TestTrie(t *testing.T) {
	testCase0(t)
	testCase1(t)
	testCase2(t)
}
