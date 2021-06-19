package Num_494_FindTargetSumWays

func findTargetSumWays(nums []int, target int) (ans int) {
	n := len(nums)
	if n == 0 {
		return
	}
	if n == 1 {
		if nums[0] == target {
			ans += 1
		}
		if nums[0] == -target {
			ans += 1
		}
		return
	}
	return findTargetSumWays(nums[1:n], target+nums[0]) + findTargetSumWays(nums[1:n], target-nums[0])
}
