package Num_63_uniquePathsWithObstacles

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 0 {
		obstacleGrid[0][0] = 1
	} else {
		return 0
	}
	for i := 1; i < len(obstacleGrid); i++ {
		if obstacleGrid[i][0] == 0 {
			obstacleGrid[i][0] = obstacleGrid[i-1][0]
		} else {
			obstacleGrid[i][0] = 0
		}
	}
	for i := 0; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
				continue
			}
			if i == 0 {
				obstacleGrid[i][j] = obstacleGrid[i][j-1]
			} else {
				obstacleGrid[i][j] = obstacleGrid[i][j-1] + obstacleGrid[i-1][j]
			}
		}
	}
	return obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
