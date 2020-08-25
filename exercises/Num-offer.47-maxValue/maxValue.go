package Num_offer_47_maxValue

func maxValue(grid [][]int) int {
	for x := 1; x < len(grid[0]); x++ {
		grid[0][x] += grid[0][x-1]
	}
	for i := 1; i < len(grid); i++ {
		grid[i][0] += grid[i-1][0]
		for j := 1; j < len(grid[i]); j++ {
			grid[i][j] += max(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
