package Num_343_integerBreak

var cache []int

func integerBreak(n int) int {
	if cache == nil {
		cache = make([]int, 59)
		cache[1] = 1
		cache[2] = 1
	} else if cache[n] != 0 {
		return cache[n]
	}
	result := 0
	for i := 1; i < n; i++ {
		result = max(result, max(integerBreak(n-i)*max(integerBreak(i), i), i*(n-i)))
	}
	cache[n] = result
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
