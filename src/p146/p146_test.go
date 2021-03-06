package p146

import (
	"testing"
)

func check(ci int, cnt *int, want, got int, t *testing.T) {
	if got != want {
		t.Errorf("Case %v [%v]: expected %v, got %v", ci, *cnt, want, got)
	}
	*cnt++
}

func testCase0(t *testing.T) {
	var cnt, got, want int
	check := func() {
		check(0, &cnt, want, got, t)
	}

	cache := Constructor(2)
	cache.Put(1,1)
	cache.Put(2,2)
	got, want = cache.Get(1),  1; check()
	cache.Put(3,3)
	got, want = cache.Get(2), -1; check()
	cache.Put(4,4)
	got, want = cache.Get(1), -1; check()
	got, want = cache.Get(3),  3; check()
	got, want = cache.Get(4),  4; check()
}

func testCase1(t *testing.T) {
	var cnt, got, want int
	check := func() {
		check(1, &cnt, want, got, t)
	}

	cache := Constructor(2)
	got, want = cache.Get(2), -1; check()
	cache.Put(2,6)
	got, want = cache.Get(1), -1; check()
	cache.Put(1,5)
	cache.Put(1,2)
	got, want = cache.Get(1),  2; check()
	got, want = cache.Get(2),  6; check()
}

func testCase2(t *testing.T) {
	var cnt, got, want int
	check := func() {
		check(2, &cnt, want, got, t)
	}

	cache := Constructor(2)
	cache.Put(2,1)
	cache.Put(3,2)
	got, want = cache.Get(3),  2; check()
	got, want = cache.Get(2),  1; check()
	cache.Put(4,3)
	got, want = cache.Get(2),  1; check()
	got, want = cache.Get(3), -1; check()
	got, want = cache.Get(4),  3; check()
}

func TestLRUCache(t *testing.T) {
	testCase0(t)
	testCase1(t)
	testCase2(t)
}
