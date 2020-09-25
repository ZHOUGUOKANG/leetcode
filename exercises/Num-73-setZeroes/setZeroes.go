package Num_73_SetZeroes

import "fmt"

func setZeroes(matrix [][]int) {
	println("start")
	printM(matrix)
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}
	printM(matrix)
	for i := 1; i < n; i++ {
		if matrix[0][i] == 0 {
			for j := 0; j < m; j++ {
				matrix[j][i] = 0
			}
		}
	}
	printM(matrix)
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			for j := 0; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}
	printM(matrix)
	if matrix[0][0] == 0 {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
	printM(matrix)
}

func printM(m [][]int) {
	for _, v := range m {
		println(fmt.Sprintf("%v", v))
	}
	println("\n")
}
