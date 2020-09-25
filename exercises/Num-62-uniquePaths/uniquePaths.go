package Num_62_UniquePaths

func uniquePaths(m int, n int) int {
	dp := make([]int, m)
	for i := 0; i < m; i++ {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		dp[0] = dp[0]
		for j := 1; j < m; j++ {
			dp[j] = dp[j-1] + dp[j]
		}
	}
	return dp[m-1]
}

//func uniquePaths(m int, n int) int {
//	visited := make([][]bool, n)
//	for i := 0; i < n; i++ {
//		visited[i] = make([]bool, m)
//	}
//	count := 0
//	next := func(i, j int) {}
//	next = func(i, j int) {
//		if i < 0 || i >= n || j < 0 || j >= m || visited[i][j] == true {
//			return
//		}
//		if i == n-1 && j == m-1 {
//			count++
//		}
//		visited[i][j] = true
//		next(i+1, j)
//		next(i, j+1)
//		visited[i][j] = false
//	}
//	next(0, 0)
//	return count
//}
