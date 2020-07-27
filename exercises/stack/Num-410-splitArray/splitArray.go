package Num_410_splitArray

import "math"

func splitArray(nums []int, m int) int {
	n := len(nums)
	f := make([][]int, n+1)
	sub := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		f[i] = make([]int, m+1)
		for j := 0; j < m+1; j++ {
			f[i][j] = math.MaxInt32
		}
	}
	for i := 0; i < n; i++ {
		sub[i+1] = sub[i] + nums[i]
	}
	f[0][0] = 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= min(i, m); j++ {
			for k := 0; k < i; k++ {
				f[i][j] = min(f[i][j], max(f[k][j-1], sub[i]-sub[k]))
			}
		}
	}
	return f[n][m]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
