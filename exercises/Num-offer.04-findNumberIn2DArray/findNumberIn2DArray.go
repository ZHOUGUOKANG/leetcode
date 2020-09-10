package Num_offer_04_findNumberIn2DArray

func findNumberIn2DArray(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}
	if matrix[0][0] > target || target > matrix[m-1][n-1] {
		return false
	}
	for i := 0; i < m; i++ {
		if matrix[i][n-1] >= target {
			for l, r := 0, n-1; l <= r; {
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
		}
	}
	return false
}
