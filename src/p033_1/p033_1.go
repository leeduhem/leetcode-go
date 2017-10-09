package p033_1

// Looks like this linear searching is fast enough for leetcode
func search(nums []int, target int) int {
	for i,v := range nums {
		if v == target {
			return i
		}
	}

	return -1
}
