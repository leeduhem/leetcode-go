package p053_1

func max(x, y int) int {
	if x > y { return x }
	return y
}

func maxSubArray(A []int) int {
	maxSoFar, maxEndingHere := A[0], A[0]
	for i := 1; i < len(A); i++ {
		maxEndingHere = max(maxEndingHere+A[i], A[i])
		maxSoFar = max(maxSoFar, maxEndingHere)
	}
	return maxSoFar
}
