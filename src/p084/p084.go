package p084

func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	index := make([]int, 0)

	maxArea := 0
	for i, height := range heights {
		for len(index) > 0 && heights[index[len(index)-1]] > height {
			h := heights[index[len(index)-1]]
			index = index[: len(index)-1]

			sidx := -1
			if len(index) > 0 {
				sidx = index[len(index)-1]
			}
			if h * (i-sidx-1) > maxArea {
				maxArea = h * (i-sidx-1)
			}
		}
		index = append(index, i)
	}

	return maxArea
}
