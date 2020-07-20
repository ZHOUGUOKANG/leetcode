package Num_16_11_divingBoard

func divingBoard(shorter int, longer int, k int) []int {
	if k == 0 {
		return []int{}
	}
	if shorter == longer {
		return []int{shorter * k}
	}
	r := make([]int, k+1)
	r[0] = shorter * k
	c := longer - shorter
	for i := 1; i <= k; i++ {
		r[i] = r[i-1] + c
	}
	return r
}
