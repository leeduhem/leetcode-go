package p120

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	minlen := make([]int, len(triangle[n-1]))
	copy(minlen, triangle[n-1])

	for layer := n-2; layer >= 0; layer-- {
		for i := 0; i <= layer; i++ {
			minlen[i] = min(minlen[i], minlen[i+1]) + triangle[layer][i]
		}
	}
	return minlen[0]
}
