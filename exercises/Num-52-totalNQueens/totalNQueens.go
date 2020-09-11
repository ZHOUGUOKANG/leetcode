package Num_52_totalNQueens

func totalNQueens(n int) (ans int) {
	queues := make([]uint16, n)
	helper := func(row int) {}
	helper = func(row int) {
		for i := 0; i < n; i++ {
			c := uint16(1 << i)
			flag := true
			for j := row - 1; j >= 0; j-- {
				jNum := queues[j]
				if jNum == c || jNum == uint16(c<<(row-j)) || jNum == uint16(c>>(row-j)) {
					flag = false
					break
				}
			}
			if flag {
				queues[row] = c
				if row == n-1 {
					ans++
				} else {
					helper(row + 1)
				}
			}
		}
	}
	helper(0)
	return
}
