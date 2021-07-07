package Num_17_NumWays

func numWays(n int, relation [][]int, k int) (ans int) {
	mapping := make(map[int][]int, n)
	for _, item := range relation {
		mapping[item[0]] = append(mapping[item[0]], item[1])
	}
	var helper func(curN, curK int)
	helper = func(curN, curK int) {
		if curK == k {
			if curN == n-1 {
				ans++
				return
			}
			return
		}
		for _, next := range mapping[curN] {
			helper(next, curK+1)
		}
	}
	helper(0, 0)
	return
}
