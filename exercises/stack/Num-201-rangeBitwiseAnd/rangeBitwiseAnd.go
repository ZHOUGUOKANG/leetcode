package Num_201_rangeBitwiseAnd

func rangeBitwiseAnd(m int, n int) int {
	shift := 0
	//求最长前缀，然后后补 固定位数的0
	for m < n {
		m, n = m>>1, n>>1
		shift++
	}
	return m << shift
}
