package Num_54_spiralOrder

const (
	LEFT = iota
	RIGHT
	UP
	DOWN
)

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])
	ans := make([]int, m*n)
	index := -1
	minM, maxM, minN, maxN := -1, m, -1, n
	helper := func(i, j, d int) {}
	helper = func(x, y, d int) {
		if x <= minM || x >= maxM || y <= minN || y >= maxN {
			return
		}
		switch d {
		case LEFT:
			for ; y > minN; y-- {
				index++
				ans[index] = matrix[x][y]
			}
			maxM--
			helper(x-1, minN+1, UP)
		case RIGHT:
			for ; y < maxN; y++ {
				index++
				ans[index] = matrix[x][y]
			}
			minM++
			helper(x+1, maxN-1, DOWN)
		case UP:
			for ; x > minM; x-- {
				index++
				ans[index] = matrix[x][y]
			}
			minN++
			helper(minM+1, y+1, RIGHT)
		case DOWN:
			for ; x < maxM; x++ {
				index++
				ans[index] = matrix[x][y]
			}
			maxN--
			helper(maxM-1, y-1, LEFT)
		}
	}
	helper(0, 0, RIGHT)
	return ans
}
