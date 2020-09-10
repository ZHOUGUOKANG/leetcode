package Num_47_permuteUnique

import (
	"math"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	visited := make([]int, len(nums))
	var result [][]int
	var ans []int
	next := func() {}
	next = func() {
		lastCase := math.MinInt32
		for k, num := range nums {
			if visited[k] == 0 {
				if lastCase == num {
					continue
				}
				lastCase = num
				visited[k] = 1
				ans = append(ans, num)
				next()
				if len(ans) == n {
					t := make([]int, n)
					copy(t, ans)
					result = append(result, t)
				}
				ans = ans[:len(ans)-1]
				visited[k] = 0
			}
		}
	}
	next()
	return result
}
