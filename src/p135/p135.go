package p135

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func candy(ratings []int) int {
	n := len(ratings)
	if (n <= 1) {
		return n
	}

	num := make([]int, n)
	for i, _ := range num {
		num[i] = 1
	}

	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			num[i] = num[i-1] + 1
		}
	}

	for i := n-1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			num[i-1] = max(num[i]+1, num[i-1])
		}
	}

	res := 0
	for _, c := range num {
		res += c
	}

	return res
}
