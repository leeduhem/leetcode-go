package p128

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func longestConsecutive(nums []int) int {
	maxLen := 0
	seqMap := map[int]int {}

	for _, n := range nums {
		if _, ok := seqMap[n]; ok {
			continue
		}

		left, right := 0, 0
		if l, ok := seqMap[n-1]; ok {
			left = l
		}
		if r, ok := seqMap[n+1]; ok {
			right = r
		}

		sum := left + right + 1
		seqMap[n] = sum

		maxLen = max(maxLen, sum)

		seqMap[n-left] = sum
		seqMap[n+right] = sum
	}

	return maxLen
}
