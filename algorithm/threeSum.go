package leetcodes

import "fmt"

/**
15. 三数之和
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

//todo 未完成
*/
func ThreeSum(nums []int) [][]int {
	result := make([][]int, 0)
	if len(nums) < 3 {
		return result
	}
	tmpResult := make(map[string]int)
	QuickSort(nums)
	min, slideMin := 0, 0
	length := len(nums)
	max := length - 1
	for {
		if nums[min] > 0 {
			break
		}
		slideMin = min + 1
		if slideMin > length-2 || nums[min]+nums[slideMin] > 0 {
			break
		}
		//if nums[min] == nums[slideMin]{
		//	slideMin++
		//	continue
		//}
		for ; slideMin < length-1; slideMin++ {
			cha := nums[min] + nums[slideMin]
			for ; max > slideMin; max-- {
				he := cha + nums[max]
				if he == 0 {
					key := fmt.Sprintf("%d~%d~%d", nums[min], nums[slideMin], nums[max])
					if _, ok := tmpResult[key]; !ok {
						tmpResult[key] = 1
						result = append(result, []int{nums[min], nums[slideMin], nums[max]})
					}
				} else if he < 0 {
					break
				}
			}
			max = length - 1
		}
		min++
		max = length - 1
	}
	return result
}
func QuickSort(s []int) []int {
	return quickSort(s, 0, len(s)-1)
}
func quickSort(s []int, left, right int) []int {
	key := s[left]
	p := left
	i, j := left, right
	for i <= j {
		for j >= p && s[j] >= key {
			j--
		}
		if j >= p {
			s[p] = s[j]
			p = j
		}

		for i <= p && s[i] <= key {
			i++
		}

		if i <= p {
			s[p] = s[i]
			p = i
		}
	}
	s[p] = key
	if p-left > 1 {
		quickSort(s, left, p-1)
	}
	if right-p > 1 {
		quickSort(s, p+1, right)
	}
	return s
}
