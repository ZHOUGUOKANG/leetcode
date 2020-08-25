package Num_39_combinationSum

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	helper := func(target, index int) [][]int { return nil }
	helper = func(target, index int) [][]int {
		var t [][]int
		for i := index; i < len(candidates) && candidates[i] <= target; i++ {
			c := candidates[i]
			if c == target {
				t = append(t, []int{c})
			} else {
				tmp := helper(target-c, i)
				if tmp == nil {
					continue
				}
				for j := 0; j < len(tmp); j++ {
					t = append(t, append([]int{c}, tmp[j]...))
				}
			}
		}
		return t
	}
	return helper(target, 0)
}
