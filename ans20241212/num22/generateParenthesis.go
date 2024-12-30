package num22

func generateParenthesis(n int) (ans []string) {
	var buf []byte
	var dfs func(l, r int)
	dfs = func(l, r int) {
		if l > n || r > n || l < r {
			return
		}
		if l == n && r == n {
			ans = append(ans, string(buf))
			return
		}
		buf = append(buf, '(')
		dfs(l+1, r)
		buf = buf[:len(buf)-1]
		buf = append(buf, ')')
		dfs(l, r+1)
		buf = buf[:len(buf)-1]
	}
	dfs(0, 0)
	return
}
