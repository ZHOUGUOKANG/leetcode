package Num_37_solveSudoku

func solveSudoku(board [][]byte) {
	row := make([]uint16, 10)
	col := make([]uint16, 10)
	box := make([]uint16, 10)
	helper := func(k, z int) bool { return true }
	next := func(i, j int) bool {
		if i == 8 && j == 8 {
			return true
		}
		if j == 8 {
			return helper(i+1, 0)
		} else {
			return helper(i, j+1)
		}
	}
	helper = func(i, j int) bool {
		c := board[i][j]
		boxNum := (i/3)*3 + j/3
		if c == '.' {
			for x := 1; x < 10; x++ {
				cur := uint16(1 << x)
				if (row[i]&cur)|(col[j]&cur)|(box[boxNum]&cur) == 0 {
					t1, t2, t3 := row[i], col[j], box[boxNum]
					row[i], col[j], box[boxNum] = row[i]|cur, col[j]|cur, box[boxNum]|cur
					if next(i, j) {
						board[i][j] = byte(x + '0')
						return true
					} else {
						row[i], col[j], box[boxNum] = t1, t2, t3
					}
				}
			}
			return false
		}
		return next(i, j)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				boxNum := (i/3)*3 + j/3
				cur := uint16(1 << (board[i][j] - '0'))
				row[i], col[j], box[boxNum] = row[i]|cur, col[j]|cur, box[boxNum]|cur
			}
		}
	}
	helper(0, 0)
}
