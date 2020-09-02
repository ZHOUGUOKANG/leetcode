package Num_46_permute

func permute(nums []int) [][]int {
	n := len(nums)
	if n == 1 {
		return [][]int{nums}
	}
	var used = make(map[int]bool, n)
	var ans [][]int
	var path []int
	dfs := func(path []int) {}
	dfs = func(path []int) {
		if len(path) == n {
			r := make([]int, n)
			copy(r, path)
			ans = append(ans, r)
			return
		}
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			dfs(path)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	dfs(path)
	return ans
}
