package Num_96_numTrees

var cache = make(map[int]int)

func numTrees(n int) int {
	if n < 2 {
		return 1
	}
	if value, ok := cache[n]; ok {
		return value
	}
	cache[0] = 1
	cache[1] = 1
	result := make([]int, n+1)
	result[0] = 1
	result[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			result[i] += result[j-1] * result[i-j]
		}
		cache[i] = result[i]
	}
	return result[n]
}
