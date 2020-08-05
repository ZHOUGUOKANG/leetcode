package Num_219_containsNearbyDuplicate

//func containsNearbyDuplicate(nums []int, k int) bool {
//	for i := 0; i < len(nums); i++ {
//		for j := i + 1; j < len(nums) && j <= i+k; j++ {
//			if nums[i] == nums[j] {
//				return true
//			}
//		}
//	}
//	return false
//}

func containsNearbyDuplicate(nums []int, k int) bool {
	k++
	cache := make(map[int]bool, k)
	for i := 0; i < len(nums); i++ {
		if len(cache) == k {
			delete(cache, nums[i-k])
		}
		if _, ok := cache[nums[i]]; ok {
			return true
		}
		cache[nums[i]] = true
	}
	return false
}
