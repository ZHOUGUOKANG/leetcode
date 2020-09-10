package Num_51_solveNQueens

import (
	"strings"
)

func solveNQueens(n int) [][]string {
	cache := make([]int, n)
	for i := 0; i < n; i++ {
		cache[i] = -1
	}
	var cacheResult [][]int
	nQueens := func(l int) {}
	nQueens = func(l int) {
		//l 代表行数，r代表列
		for r := 0; r < n; r++ {
			flag := true
			for j, left, right := l-1, r-1, r+1; j >= 0; j, left, right = j-1, left-1, right+1 {
				cj := cache[j]
				if cj == r || (left >= 0 && cj == left) || (right < n && cj == right) {
					flag = false
					break
				}
			}
			if !flag {
				continue
			}
			cache[l] = r
			if l == n-1 {
				t := make([]int, n)
				copy(t, cache)
				cacheResult = append(cacheResult, t)
			} else {
				nQueens(l + 1)
			}
		}
	}
	nQueens(0)
	result := make([][]string, len(cacheResult))
	if len(cacheResult) > 0 {
		template := []byte(strings.Repeat(".", n))
		for i := 0; i < len(cacheResult); i++ {
			z := cacheResult[i]
			c := make([]string, n)
			for j := 0; j < n; j++ {
				template[z[j]] = 'Q'
				c[j] = string(template)
				template[z[j]] = '.'
			}
			result[i] = c
		}
	}
	return result
}
