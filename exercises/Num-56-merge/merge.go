package Num_56_Merge

//
//func merge(intervals [][]int) [][]int {
//	n := len(intervals)
//	if n == 0 {
//		return nil
//	}
//	var ans = [][]int{intervals[0]}
//	has := false
//	for i := 1; i < n; i++ {
//		c := intervals[i]
//		flag := true
//		for j := 0; j < len(ans); j++ {
//			a := ans[j]
//			if c[0] > a[1] || c[1] < a[0] {
//				continue
//			}
//			has =true
//			a[0] = min(a[0],c[0])
//			a[1] = max(a[1],c[1])
//			flag = false
//			break
//		}
//		if flag {
//			ans = append(ans, c)
//		}
//	}
//	if has {
//		return merge(ans)
//	}else {
//		return ans
//	}
//
//}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	left, right := 0, 1
	changed := false
	for right < len(intervals) {
		merged := false
		for i := 0; i <= left; i++ {
			if intervals[right][0] > intervals[i][1] || intervals[right][1] < intervals[i][0] {
				continue
			}
			intervals[i][0] = min(intervals[right][0], intervals[i][0])
			intervals[i][1] = max(intervals[right][1], intervals[i][1])
			changed = true
			merged = true
			break
		}
		if !merged {
			left++
			intervals[left] = intervals[right]
		}
		right++
	}
	if changed {
		return merge(intervals[:left+1])
	} else {
		return intervals[:left+1]
	}
}
