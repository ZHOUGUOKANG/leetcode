package Num_378_kthSmallest

import "math"

func kthSmallest(matrix [][]int, k int) int {
	length := len(matrix)
	cache := make([]int, len(matrix))
	lastValue := 0
	for i := 0; i < k; i++ {
		tmpIndex := -1
		tmpValue := math.MaxInt32
		for j := 0; j < len(cache); j++ {
			if cache[j] < length {
				if matrix[j][cache[j]] < tmpValue {
					tmpValue = matrix[j][cache[j]]
					tmpIndex = j
					lastValue = tmpValue
				}
			}
		}
		cache[tmpIndex]++
	}
	return lastValue
}
