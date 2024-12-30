package num16

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) (ans int) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	n := len(nums)
	minAbs := math.MaxFloat64
	for i := 0; i < n-2; i++ {
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if mm := math.Abs(float64(sum - target)); mm < minAbs {
				ans = sum
				minAbs = mm
			}
			if sum < target {
				j++
			} else if sum > target {
				k--
			} else {
				j, k = j+1, k-1
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}
	return
}
