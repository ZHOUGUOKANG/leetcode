package Num_44_isMatch

func isMatch(s string, p string) bool {
	n, m := len(s), len(p)
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, m+1)
	}
	for i := 1; i <= m; i++ {
		if string(p[i-1]) == "*" {
			dp[0][i] = true
		} else {
			break
		}
	}
	dp[0][0] = true
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if p[j-1] == 42 {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else if p[j-1] == 63 || p[j-1] == s[i-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[n][m]
}
