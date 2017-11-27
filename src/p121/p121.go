package p121

func maxProfit(prices []int) int {
	profit := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j] - prices[i] > profit {
				profit = prices[j] - prices[i]
			}
		}
	}

	return profit
}
