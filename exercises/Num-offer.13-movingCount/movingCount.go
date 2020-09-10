package Num_offer_13_movingCount

func movingCount(m int, n int, k int) int {
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	count := 0
	helper := func(a, b int) {}
	helper = func(a, b int) {
		if a < 0 || b < 0 || a >= m || b >= n || visited[a][b] == false {
			return
		}
		c := 0
		x, y := a, b
		for x != 0 || y != 0 {
			c += x%10 + y%10
			x, y = x/10, y/10
		}
		if c <= k {
			visited[a][b] = true
			count++
			helper(a-1, b)
			helper(a+1, b)
			helper(a, b-1)
			helper(a, b+1)
		}
	}
	helper(0, 0)
	return count
}
