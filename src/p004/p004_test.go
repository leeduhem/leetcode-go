package p004

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	nums11 := []int {1,3}
	nums12 := []int {2}
	med1   := 2.0
	res1   := findMedianSortedArrays(nums11, nums12)

	if (med1 != res1) {
		t.Errorf("Expected %v, got %v", med1, res1)
	}

	nums21 := []int {1,2}
	nums22 := []int {3,4}
	med2   := (2+3)/2.0
	res2   := findMedianSortedArrays(nums21, nums22)

	if (med2 != res2) {
		t.Errorf("Expected %v, got %v", med2, res2)
	}
}
