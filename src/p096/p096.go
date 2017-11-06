package p096

func numTrees(n int) int {
	tn := make([]int, n+1)

	var numberOfTrees func (n int) int;
	numberOfTrees = func (n int) int {
		if tn[n] != 0 {
			return tn[n]
		}

		if n == 0 {
			tn[n] = 1
			return 1
		}

		cnt := 0
		for i := 0; i < n; i++ {
			cnt += numberOfTrees(i) * numberOfTrees(n-i-1)
		}

		tn[n] = cnt
		return cnt
	}

	return numberOfTrees(n)
}
