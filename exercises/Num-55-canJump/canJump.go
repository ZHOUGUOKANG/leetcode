package Num_55_canJump

func canJump(nums []int) bool {
	n := len(nums)
	max := 0
	for i := 0; i <= max && i < n; i++ {
		c := nums[i] + i
		if c > max {
			max = c
		}
	}
	return max >= n-1
}
