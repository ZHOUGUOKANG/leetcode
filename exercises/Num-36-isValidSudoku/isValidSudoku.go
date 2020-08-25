package Num_36_isValidSudoku

func isValidSudoku(board [][]byte) bool {
	var row, col, box [9]uint16
	var cur uint16
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			cur = 1 << (board[i][j] - '0')
			boxIndex := j/3 + (i/3)*3
			if (row[i]&cur)|(col[j]&cur)|(box[boxIndex]&cur) != 0 {
				return false
			}
			row[i], col[j], box[boxIndex] = cur, cur, cur
		}
	}
	return true
}
