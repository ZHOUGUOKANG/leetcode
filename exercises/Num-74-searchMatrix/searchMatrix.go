package Num_74_SearchMatrix

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}
	l, r := 0, m-1
	for l <= r {
		i := (l + r) / 2
		last := matrix[i][n-1]
		if l == r {
			l, r = 0, n-1
			for l <= r {
				mid := (l + r) / 2
				c := matrix[i][mid]
				if c == target {
					return true
				} else if c > target {
					r = mid - 1
				} else {
					l = mid + 1
				}
			}
		} else if last >= target {
			r = i
		} else if last < target {
			l = i + 1
		}
	}
	return false
}
