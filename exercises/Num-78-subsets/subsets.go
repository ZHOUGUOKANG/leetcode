package Num_78_Subsets

//func subsets(nums []int) [][]int {
//	if len(nums) == 1{
//		return [][]int{{},{nums[0]}}
//	}
//	if len(nums) == 0{
//		return [][]int{{}}
//	}
//	var ans [][]int
//	for i:=0;i<len(nums);i++{
//		sets := subsets(nums[i+1:])
//		for _,v := range sets{
//			ans = append(ans,append([]int{nums[i]},v...))
//		}
//	}
//	ans = append(ans,[]int{})
//	return ans
//}

func subsets(nums []int) (ans [][]int) {
	var set []int
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			t := make([]int, len(set))
			copy(t, set)
			ans = append(ans, t)
			return
		}
		//选中
		set = append(set, nums[cur])
		dfs(cur + 1)
		//不选中
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return
}
