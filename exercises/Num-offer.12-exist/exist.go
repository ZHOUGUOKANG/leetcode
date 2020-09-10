package Num_offer_12_exist

func exist(board [][]byte, word string) bool {
	w := []byte(word)
	m, n, wl := len(board), len(board[0]), len(w)
	if m*n < wl || wl < 1 {
		return false
	}
	next := func(i, j, wi int) bool { return false }
	next = func(i, j, wi int) bool {
		if wi == wl {
			return true
		}
		if i < 0 || i >= m || j < 0 || j >= n {
			return false
		}
		if board[i][j] != w[wi] {
			return false
		}
		t := board[i][j]
		board[i][j] = '.'
		if next(i-1, j, wi+1) || next(i+1, j, wi+1) || next(i, j-1, wi+1) || next(i, j+1, wi+1) {
			return true
		}
		board[i][j] = t
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == w[0] {
				if next(i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}
