package Num_16_threeSumClosest

import (
	"math"
	"sort"
)

//func threeSumClosest(nums []int, target int) int {
//	bestCha := math.MaxInt32
//	best := 0
//	for i:=0;i<len(nums)-2;i++{
//		for j:=i+1;j<len(nums)-1;j++{
//			for x:=j+1;x<len(nums);x++{
//				cha := abs(nums[i] + nums[j] + nums[x] - target)
//				if cha < bestCha{
//					bestCha = cha
//					best = nums[i] + nums[j] + nums[x]
//				}
//			}
//		}
//	}
//	return best
//}
//
//func abs(a int)int  {
//	if a >= 0{
//		return a
//	}
//	return -a
//}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	bestCha := math.MaxInt32
	best := 0
	for x := 0; x < len(nums)-2; x++ {
		y := x + 1
		z := len(nums) - 1
		for y < z {
			sum := nums[x] + nums[y] + nums[z]
			cha := abs(sum - target)
			if sum > target {
				z--
			} else if sum < target {
				y++
			} else {
				return target
			}
			if cha < bestCha {
				bestCha = cha
				best = sum
			}
		}
	}
	return best
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
