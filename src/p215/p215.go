package p215

import (
	"sort"
)

func findKthLargest(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[j] < nums[i]
	})
	return nums[k-1]
}
