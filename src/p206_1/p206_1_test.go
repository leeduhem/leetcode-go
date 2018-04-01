package p206_1

import (
	"testing"
)

func TestReverseList(t *testing.T) {
	testCase := func(in, want *ListNode) {
		in0 := deepCopyList(in)
		got := reverseList(in)
		if !isListEqual(got, want) {
			t.Errorf("Case %v: expected %v, got %v", in0, want, got)
		}
	}

	testCase(ints2List([]int{1, 2, 3}), ints2List([]int{3, 2, 1}))
}
