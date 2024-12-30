package num18

import "sort"

func fourSum(nums []int, target int) (ans [][]int) {
	if len(nums) < 4 {
		return
	}
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			k, l := j+1, n-1
			for k < l {
				sum := nums[i] + nums[j] + nums[k] + nums[l]
				if sum < target {
					k++
				} else if sum > target {
					l--
				} else {
					ans = append(ans, []int{nums[i], nums[j], nums[k], nums[l]})
					k, l = k+1, l-1
					for k < l && nums[k] == nums[k-1] {
						k++
					}
					for k < l && nums[l] == nums[l+1] {
						l--
					}
				}
			}
		}
	}
	return
}
