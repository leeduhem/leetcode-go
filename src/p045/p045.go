package p045

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func jump(nums []int) int {
	nlen := len(nums)
	if nlen < 2 {
		return 0
	}

	// breadth-first searching
	level, i, currentMax, nextMax := 0, 0, 0, 0
	for currentMax-i+1 > 0 { // nodes of current level > 0
		level++

		// travel current level, and update the max reachable from next level
		for ; i <= currentMax; i++ {
			nextMax = max(nextMax, nums[i]+i)

			// last element is in level+1, so the min jump is level
			if nextMax >= nlen-1 {
				return level
			}
		}
		currentMax = nextMax
	}

	return 0
}
