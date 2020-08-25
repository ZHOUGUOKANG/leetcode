package Num_121_MaxProfit

func MaxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	min := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > min {
			tmpProfit := prices[i] - min
			if tmpProfit > maxProfit {
				maxProfit = tmpProfit
			}
		} else {
			min = prices[i]
		}
	}
	return maxProfit
}
