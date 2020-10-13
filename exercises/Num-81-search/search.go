package Num_81_Search

func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		midNum := nums[mid]
		if midNum == target {
			return true
		}
		if midNum == nums[l] {
			l++
		} else if midNum == nums[r] {
			r--
		} else if midNum < nums[r] {
			if target > midNum && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else if midNum > nums[l] {
			if target < midNum && target >= nums[l] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return false
}
