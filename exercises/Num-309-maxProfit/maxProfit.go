package Num_309_maxProfit

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	var r0, r1, r2 = -prices[0], 0, 0
	p := 0
	for i := 1; i < len(prices); i++ {
		p = prices[i]
		r0, r1, r2 = maxInt(r2-p, r0), r0+p, maxInt(r1, r2)
	}
	return maxInt(r1, r2)
}

func maxInt(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
