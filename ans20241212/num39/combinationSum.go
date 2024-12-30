package num39

import "sort"

func combinationSum(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	cur := make([]int, 0, 16)
	var helper func(i, t int)
	helper = func(i, t int) {
		if t == 0 {
			k := make([]int, len(cur))
			copy(k, cur)
			ans = append(ans, k)
			return
		}
		if i > len(candidates) || t < candidates[i] {
			return
		}
		for j := i; j < len(candidates); j++ {
			cur = append(cur, candidates[j])
			helper(j, t-candidates[j])
			cur = cur[:len(cur)-1]
		}
	}
	helper(0, target)
	return
}
