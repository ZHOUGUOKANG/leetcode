package Num_59_GenerateMatrix

func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := 0; i < len(ans); i++ {
		ans[i] = make([]int, n)
	}
	num := 1
	l, r, u, d, t := 0, n-1, 0, n-1, 0
	for num <= n*n {
		for t = l; t <= r; t++ {
			ans[u][t] = num
			num++
		}
		u++
		for t = u; t <= d; t++ {
			ans[t][r] = num
			num++
		}
		r--
		for t = r; t >= l; t-- {
			ans[d][t] = num
			num++
		}
		d--
		for t = d; t >= u; t-- {
			ans[t][l] = num
			num++
		}
		l++
	}
	return ans
}
