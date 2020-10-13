package Num_416_CanPartition

func canPartition(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}
	sum, max := 0, -1
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}
	target := sum / 2
	if sum%2 == 1 || max > target {
		return false
	}
	dp := make([]bool, target+1)
	dp[0] = true
	for _, v := range nums {
		for j := target; j >= v; j-- {
			dp[j] = dp[j] || dp[j-v]
		}
	}
	return dp[target]
}
