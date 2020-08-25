package Num_566_matrixReshape

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums[0])*len(nums) != r*c {
		return nums
	}
	result := make([][]int, r)
	for i := 0; i < r; i++ {
		result[i] = make([]int, c)
	}
	index := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			h := index / c
			l := index - h*c
			result[h][l] = nums[i][j]
			index++
		}
	}
	return result
}
