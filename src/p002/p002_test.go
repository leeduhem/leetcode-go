package p002

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	testCase := func(in0, in1, want *ListNode) {
		got := addTwoNumbers(in0, in1)
		if !isListEqual(got, want) {
			t.Errorf("Case %v, %v: expected %v, got %v",
				in0, in1, want, got)
		}
	}

	in00 := ints2List([]int{2, 4, 3})
	in01 := ints2List([]int{5, 6, 4})
	want0 := ints2List([]int{7, 0, 8})
	testCase(in00, in01, want0)

	in10 := ints2List([]int{5})
	want1 := ints2List([]int{0, 1})
	testCase(in10, in10, want1)
}
