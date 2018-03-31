package p203

import (
	"testing"
)

func TestRemoveElements(t *testing.T) {
	testCase := func(in0 *ListNode, in1 int, want *ListNode) {
		in00 := deepCopyList(in0)
		got := removeElements(in0, in1)
		if !isListEqual(got, want) {
			t.Errorf("Case %v, %v: expected %v, got %v", in00, in1, want, got)
		}
	}

	in00 := ints2List([]int{1, 2, 6, 3, 4, 5, 6})
	want0 := ints2List([]int{1, 2, 3, 4, 5})
	testCase(in00, 6, want0)
}
