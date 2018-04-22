package p004

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	var tests = []struct {
		in0, in1 []int
		want     float64
	}{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
	}

	for _, test := range tests {
		got := findMedianSortedArrays(test.in0, test.in1)
		if got != test.want {
			t.Errorf("findMedianSortedArrays(%#v, %#v) = %v, want %v",
				test.in0, test.in1, got, test.want)
		}
	}
}
