package p035

import "sort"

// cheating
func searchInsert(nums []int, target int) int {
    return sort.SearchInts(nums, target)
}
