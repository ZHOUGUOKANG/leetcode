package Num_532_FindPairs

//import "sort"
//
//func findPairs(nums []int, k int) int {
//	if k < 0{
//		return 0
//	}
//	sort.Ints(nums)
//	count, n := 0, len(nums)
//	for i := 0; i < n-1; i++ {
//		if i > 0 && nums[i] == nums[i-1] {
//			continue
//		}
//		for j := i + 1; j < n; j++ {
//			if j > i+1 && nums[j] == nums[j-1] {
//				continue
//			}
//			if nums[j]-nums[i] == k {
//				count++
//			}
//		}
//	}
//	return count
//}

func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	numsHas := make(map[int]bool)
	diffHas := make(map[int]bool)

	for _, num := range nums {
		if numsHas[num-k] {
			diffHas[num-k] = true
		}
		if numsHas[num+k] {
			diffHas[num+k] = true
		}
		numsHas[num] = true
	}
	return len(diffHas)
}
