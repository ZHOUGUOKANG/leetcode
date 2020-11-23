package Num_327_CountRangeSum

func countRangeSum(nums []int, lower int, upper int) (ans int) {
	for i, c, n := 0, 0, len(nums); i < n; i++ {
		c = 0
		for j := i; j < n; j++ {
			c += nums[j]
			if c >= lower && c <= upper {
				ans++
			}
		}
	}
	return
}
