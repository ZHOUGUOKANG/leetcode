package Num_714_maxProfit

func maxProfit(prices []int, fee int) int {
	cash, hold := 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		cash, hold = max(cash, hold+prices[i]-fee), max(hold, cash-prices[i])
	}
	return max(cash, hold)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
