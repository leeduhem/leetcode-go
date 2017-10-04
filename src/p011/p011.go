package p011

func min(x,y int) int {
	if x < y {
		return x
	}
	return y
}

func maxArea(height []int) int {
	max := 0

	for i,ai := range height {
		for j,aj := range height[i+1:] {
			// (j + i+1) - i = j + 1
			area := min(ai, aj) * (j + 1)
			if (area > max) {
				max = area
			}
		}
	}

	return max
}
