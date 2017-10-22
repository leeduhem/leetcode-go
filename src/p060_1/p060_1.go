package p060_1

func getPermutation(n, k int) string {
	if n > 9 {
		panic("Invalid input")
	}

	factorial := make([]int, n+1)
	numbers   := make([]int, n)
	str       := make([]byte, n)

	factorial[0] = 1
	for sum, i := 1, 1; i <= n; i++ {
		sum *= i
		factorial[i] = sum
		numbers[i-1] = i
	}

	k--

	for i := 1; i <= n; i++ {
		i1 := k / factorial[n-i]
		str[i-1] = "0123456789"[numbers[i1]]
		numbers = append(numbers[:i1], numbers[i1+1:]...)
		k -= i1 * factorial[n-i]
	}

	return string(str)
}
