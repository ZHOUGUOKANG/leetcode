package Num_77_combine

import "fmt"

func combine(n int, k int) [][]int {
	var result [][]int
	if k == 0 || n < 0 {
		return result
	}
	cache := map[string][][]int{}
	helper := func(x, y int) [][]int { return nil }
	helper = func(x, y int) [][]int {
		cacheKey := fmt.Sprintf("%d-%d", x, y)
		if v, ok := cache[cacheKey]; ok {
			return v
		}
		var r [][]int
		if y == 1 {
			r = make([][]int, 0)
			for ; x <= n; x++ {
				r = append(r, []int{x})
			}
		} else {
			for i := x + 1; i <= n; i++ {
				t := helper(i, y-1)
				for j := 0; j < len(t); j++ {
					a := make([]int, y)
					a[0] = i - 1
					copy(a[1:], t[j])
					r = append(r, a)
				}
			}
		}
		cache[cacheKey] = r
		return r
	}
	return helper(1, k)
}
