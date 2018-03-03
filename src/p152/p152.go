package p152

func intMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func intMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxProduct(nums []int) int {
	r := nums[0]

	for i, imax, imin := 1, r, r; i < len(nums); i++ {
		if nums[i] < 0 {
			imax, imin = imin, imax
		}

		imax = intMax(nums[i], imax * nums[i])
		imin = intMin(nums[i], imin * nums[i])

		r = intMax(r, imax)
	}

	return r
}
