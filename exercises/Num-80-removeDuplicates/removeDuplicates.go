package Num_80_RemoveDuplicates

func removeDuplicates(nums []int) int {
	l, r, count := 0, 1, 1
	for ; r < len(nums); r++ {
		if nums[r] == nums[l] {
			if count < 2 {
				count++
				l++
				nums[l] = nums[r]
			}
		} else {
			l++
			nums[l] = nums[r]
			count = 1
		}
	}
	return l + 1
}
