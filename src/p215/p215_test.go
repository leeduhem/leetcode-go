package p215

import (
	"testing"
	. "testing/utils"
)

func TestFindKthLargest(t *testing.T) {
	testCase := func(in0 []int, in1, want int) {
		in00 := DeepCopyInts(in0)
		if got := findKthLargest(in0, in1); got != want {
			t.Errorf("Case %v, %v: expected %v, got %v",
				in00, in1, want, got)
		}
	}

	testCase([]int{3, 2, 1, 5, 6, 4}, 2, 5)
}
