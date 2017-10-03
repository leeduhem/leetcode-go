package p001

import "testing"

func TestTwoSum(t *testing.T) {
	nums1 := []int {3, 2, 4}
	target1 := 6
	answer1 := []int {1, 2}

	indins1 := twoSum(nums1, target1)
	if (len(indins1) != len(answer1)) {
		t.Error("Expected ", answer1)
	}
	for i,v := range indins1 {
		if (v != answer1[i]) {
			t.Errorf("Expected %c, got %c", answer1[i], v)
		}
	}
}
