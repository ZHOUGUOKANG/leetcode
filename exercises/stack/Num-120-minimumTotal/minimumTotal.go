package Num_120_minimumTotal

func minimumTotal(triangle [][]int) int {
	length := len(triangle)
	for i := 1; i < length; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				triangle[i][j] += triangle[i-1][j]
			} else if j == len(triangle[i])-1 {
				triangle[i][j] += triangle[i-1][j-1]
			} else {
				if triangle[i-1][j] > triangle[i-1][j-1] {
					triangle[i][j] += triangle[i-1][j-1]
				} else {
					triangle[i][j] += triangle[i-1][j]
				}
			}
		}
	}
	min := triangle[length-1][0]
	for i := 1; i < len(triangle[length-1]); i++ {
		if min > triangle[length-1][i] {
			min = triangle[length-1][i]
		}
	}
	return min
}
