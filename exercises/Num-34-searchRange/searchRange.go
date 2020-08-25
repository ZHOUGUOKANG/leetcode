package Num_34_searchRange

func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 || target < nums[0] || target > nums[n-1] {
		return []int{-1, -1}
	}
	l, r, mid := 0, n-1, 0
	for l <= r {
		mid = (l + r) / 2
		midNum := nums[mid]
		if midNum == target {
			l, r = mid-1, mid+1
			for l > -1 && nums[l] == target {
				l--
			}
			for r < n && nums[r] == target {
				r++
			}
			l, r = l+1, r-1
			return []int{l, r}
		} else if midNum >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return []int{-1, -1}
}
