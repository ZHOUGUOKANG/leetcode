package Num_279_NumSquares

import (
	"math"
)

var cache = map[int]int{}

func numSquares(n int) (ans int) {
	if v, ok := cache[n]; ok {
		return v
	}
	ans = math.MaxInt64
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		cha := n - i*i
		if cha == 0 {
			cache[n] = 1
			return 1
		}
		if cha < 0 {
			return 0
		}
		if count := numSquares(cha); count > 0 {
			ans = min(1+count, ans)
		}
	}
	cache[n] = ans
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
