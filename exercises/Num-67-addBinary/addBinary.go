package Num_67_AddBinary

func addBinary(a string, b string) string {
	m, n := len(a), len(b)
	maxLen := max(m, n)
	ans := make([]byte, maxLen+1)
	carry := 0
	for i := 0; i < maxLen; i++ {
		if i < m {
			carry += int(a[m-i-1] - 48)
		}
		if i < n {
			carry += int(b[n-i-1] - 48)
		}
		ans[maxLen-i] = byte(carry%2 + 48)
		carry /= 2
	}
	if carry > 0 {
		ans[0] = '1'
		return string(ans)
	}
	return string(ans[1:])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
