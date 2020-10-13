package Num_832_FlipAndInvertImage

func flipAndInvertImage(A [][]int) [][]int {
	for i := 0; i < len(A); i++ {
		lj := len(A[i])
		mid := lj / 2
		for j := 0; j <= mid; j++ {
			A[i][j], A[i][lj-1-j] = A[i][lj-1-j], A[i][j]
			if A[i][j] == 1 {
				A[i][j] = 0
			} else {
				A[i][j] = 1
			}
			if A[i][lj-1-j] == 1 {
				A[i][lj-1-j] = 0
			} else {
				A[i][lj-1-j] = 1
			}
		}
		if lj%2 != 0 {
			if A[i][mid] == 0 {
				A[i][mid] = 1
			} else {
				A[i][mid] = 0
			}
		}

	}
	return A
}
