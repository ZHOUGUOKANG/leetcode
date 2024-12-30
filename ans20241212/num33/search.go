package num33

func search(nums []int, target int) (ans int) {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) >> 1
		m, n, k := nums[l], nums[r], nums[mid]
		if k == target {
			return mid
		}
		//k-n单调递增
		if k < n {
			if target > n || target < k {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else { //m到k单调递增
			if target > k || target < m {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
