package Num_1893_IsCovered

func isCovered(ranges [][]int, left int, right int) bool {
	var t uint64
	for _, item := range ranges {
		for i := item[0]; i < item[1]; i++ {
			t |= 1 << i
		}
	}
	for i := left; i <= right; i++ {
		if t&1<<i == 0 {
			return false
		}
	}
	return true
}
