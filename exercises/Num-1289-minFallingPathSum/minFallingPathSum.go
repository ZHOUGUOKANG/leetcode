package Num_1289_minFallingPathSum

import (
	"math"
)

func minFallingPathSum(arr [][]int) int {
	result := make([][]int, len(arr))
	for index, value := range arr {
		result[index] = make([]int, len(value))
		for k, v := range value {
			if index == 0 {
				result[index][k] = v
			} else {
				min := math.MaxInt32
				for i := 0; i < len(result[index-1]); i++ {
					if i != k && result[index-1][i] < min {
						min = result[index-1][i]
					}
				}
				result[index][k] = v + min
			}
		}
	}
	tail := len(result) - 1
	min := result[tail][0]
	for i := 1; i < len(result[tail]); i++ {
		if result[tail][i] < min {
			min = result[tail][i]
		}
	}
	return min
}
