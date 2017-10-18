package p042

func min(x,y int) int {
	if x < y {
		return x
	}
	return y
}

func sunkenArea(height []int) int {
	low := min(height[0], height[len(height)-1])

	area := 0
	for _,h := range height {
		if h > low {
			continue
		}
		area += low - h
	}

	return area
}

func findPeaks(height []int) []int {
	hlen := len(height)
	indexes := []int {}

	i := 0

	for i < hlen {
		for ; i+1 < hlen && height[i] <= height[i+1]; i++ {}
		indexes = append(indexes, i)
		if i+1 == hlen {
			return indexes
		}

		// Ignore flat ground and/or downhill
		for ; i+1 < hlen && height[i] >= height[i+1]; i++ {}
		if i+1 == hlen {
			return indexes
		}
	}

	return indexes
}

func trap(height []int) int {
	peaks := findPeaks(height)
	plen := len(peaks)

	area := 0
	for i := 0; i+1 < plen; {
		m, mj := -1, -1
		j := i+1
		for ; j < plen && height[peaks[i]] > height[peaks[j]]; j++ {
			if m < height[peaks[j]] {
				m = height[peaks[j]]
				mj = j
			}
		}

		left, right := -1, -1
		if j < plen {
			left, right = peaks[i], peaks[j]
			i = j
		} else {
			left, right = peaks[i], peaks[mj]
			i = mj
		}

		area += sunkenArea(height[left:right+1])
	}

	return area
}
