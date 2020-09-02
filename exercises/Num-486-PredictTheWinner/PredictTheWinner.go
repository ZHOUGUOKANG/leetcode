package Num_486_PredictTheWinner

func PredictTheWinner(nums []int) bool {
	length := len(nums)
	dp := make([]int, length)
	copy(dp, nums)
	for i := length - 2; i >= 0; i-- {
		for j := i + 1; j < length; j++ {
			dp[j] = max(nums[i]-dp[j], nums[j]-dp[j-1])
		}
	}
	return dp[length-1] >= 0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
