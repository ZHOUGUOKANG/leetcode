package Num_40_combinationSum2

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return helper(candidates, -1, target)
}

func helper(candidates []int, start int, target int) [][]int {
	var r [][]int
	if target == 0 {
		return r
	}
	for i := start + 1; i < len(candidates) && candidates[i] <= target; i++ {
		c := candidates[i]
		if i > start+1 && c == candidates[i-1] {
			continue
		}
		if c == target {
			r = append(r, []int{c})
		} else {
			tmp := helper(candidates, i, target-c)
			if len(tmp) != 0 {
				for j := 0; j < len(tmp); j++ {
					r = append(r, append([]int{c}, tmp[j]...))
				}
			}
		}
	}
	return r
}
