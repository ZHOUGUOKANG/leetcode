package Num_64_minPathSum

func minPathSum(grid [][]int) int {
	for x := 1; x < len(grid[0]); x++ {
		grid[0][x] += grid[0][x-1]
	}
	for i := 1; i < len(grid); i++ {
		grid[i][0] += grid[i-1][0]
		for j := 1; j < len(grid[i]); j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
