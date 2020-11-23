package Num_452_FindMinArrowShots

import (
	"sort"
)

func findMinArrowShots(points [][]int) (ans int) {
	n := len(points)
	if n <= 1 {
		return n
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	for cur, i := points[0], 1; i < len(points); i++ {
		t := points[i]
		if t[0] > cur[1] {
			ans, cur = ans+1, t
		} else {
			if cur[0] < t[0] {
				cur[0] = t[0]
			}
			if cur[1] > t[1] {
				cur[1] = t[1]
			}
		}
	}
	return ans + 1
}
