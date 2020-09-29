package Num_90_SubsetsWithDup

import (
	"sort"
)

//
//func subsetsWithDup(nums []int) [][]int {
//	sort.Ints(nums)
//	var helper func(nums []int) [][]int
//	helper = func(nums []int) [][]int {
//		ans := [][]int{{}}
//		for i := 0; i < len(nums); i++ {
//			if i > 0 && nums[i] == nums[i-1] {
//				continue
//			}
//			tNums := helper(nums[i+1:])
//			for j := 0; j < len(tNums); j++ {
//				tNums[j] = append(tNums[j], nums[i])
//				tNums[j][0], tNums[j][len(tNums[j])-1] = tNums[j][len(tNums[j])-1], tNums[j][0]
//				ans = append(ans, tNums[j])
//			}
//		}
//		return ans
//	}
//	return helper(nums)
//}

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	dfs := func(temp []int, start int) {}
	dfs = func(temp []int, start int) {
		tmp := make([]int, len(temp))
		copy(tmp, temp)
		res = append(res, tmp)
		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			temp = append(temp, nums[i])
			dfs(temp, i+1)
			temp = temp[:len(temp)-1]
		}
	}
	dfs([]int{}, 0)
	return res
}
