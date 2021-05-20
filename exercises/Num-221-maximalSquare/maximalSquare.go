package Num_221_MaximalSquare

func maximalSquare(matrix [][]byte) (ans int) {
	x := len(matrix)
	if x == 0 {
		return 0
	}
	y := len(matrix[0])
	min := x
	if y < x {
		min = y
	}
	for length := 1; length <= min; length++ {
		for startI := 0; startI <= x-length; startI++ {
			for startJ := 0; startJ <= y-length; startJ++ {
				flag := true
				for i := startI; (i < startI+length) && flag; i++ {
					for j := startJ; j < startJ+length; j++ {
						if matrix[i][j] == '0' {
							flag = false
							break
						}
					}
				}
				if flag {
					ans = length * length
				}
			}
		}
	}
	return
}
