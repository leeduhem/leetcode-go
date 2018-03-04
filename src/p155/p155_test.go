package p155

import (
	"testing"
)

func check(cn int, idx *int, want, got int, t *testing.T) {
	if got != want {
		t.Errorf("Case %v [%v]: expected %v, got %v", cn, *idx, want, got)
	}
	*idx++
}

func testCase0(t *testing.T) {
	var idx, got, want int
	check := func() {
		check(0, &idx, want, got, t)
	}

	stack := Constructor()

	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	got, want = stack.GetMin(), -3; check()
	stack.Pop()
	got, want = stack.Top(), 0; check()
	got, want = stack.GetMin(), -2; check()
}

func testCase1(t *testing.T) {
	var idx, want, got int
	check := func() {
		check(1, &idx, want, got, t)
	}

	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	got, want = stack.Top(), 2; check()
	got, want = stack.GetMin(), 1; check()
	stack.Pop()
	got, want = stack.GetMin(), 1; check()
	got, want = stack.Top(), 1; check()
}

func TestMinStack(t *testing.T) {
	testCase0(t)
	testCase1(t)
}
