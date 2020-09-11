package Num_216_combinationSum3

func combinationSum3(k int, n int) [][]int {
	var ans [][]int
	var tmp []int
	helper := func(start, num int) {}
	helper = func(start, num int) {
		if len(tmp) == k {
			if num == 0 {
				t := make([]int, k)
				copy(t, tmp)
				ans = append(ans, t)
			}
			return
		}
		if num < start || len(tmp) > k {
			return
		}
		for i := start; i <= 9 && i <= num; i++ {
			tmp = append(tmp, i)
			helper(i+1, num-i)
			tmp = tmp[:len(tmp)-1]
		}
	}
	helper(1, n)
	return ans
}
