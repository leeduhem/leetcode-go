package p188

func intMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxProfit(k int, prices []int) int {
	pn := len(prices)

	if k >= pn / 2 {
		profit := 0
		for i := 1; i < pn; i++ {
			if prices[i] > prices[i-1] {
				profit += prices[i] - prices[i-1]
			}
		}
		return profit
	}

	t := make([][]int, k+1)
	for i := range t {
		t[i] = make([]int, pn)
	}

	for i := 1; i <= k; i++ {
		tmpMax := -prices[0]
		for j := 1; j < pn; j++ {
			t[i][j] = intMax(t[i][j-1], prices[j] + tmpMax)
			tmpMax = intMax(tmpMax, t[i-1][j-1] - prices[j])
		}
	}

	return t[k][pn-1]
}
