package p216

func combination(ans *[][]int, comb []int, k, start, n int) {
	if len(comb) > k {
		return
	}

	if len(comb) == k && n == 0 {
		comb1 := make([]int, len(comb))
		copy(comb1, comb)
		*ans = append(*ans, comb1)
		return
	}

	for i := start; i <= n && i <= 9; i++ {
		comb1 := append(comb, i)
		combination(ans, comb1, k, i+1, n-i)
	}
}

func combinationSum3(k int, n int) [][]int {
	ans := [][]int{}
	combination(&ans, []int{}, k, 1, n)
	return ans
}
