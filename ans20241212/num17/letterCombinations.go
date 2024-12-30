package num17

var d = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) (ans []string) {
	if len(digits) == 0 {
		return
	}
	buf := make([]byte, 0, 4)
	var dfs func(int)
	dfs = func(i int) {
		if i >= len(digits) {
			ans = append(ans, string(buf))
			return
		}
		for _, t := range d[digits[i]-'2'] {
			buf = append(buf, byte(t))
			dfs(i + 1)
			buf = buf[:len(buf)-1]
		}
	}
	dfs(0)
	return
}
