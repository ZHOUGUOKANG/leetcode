package Num_16_05_trailingZeroes

func trailingZeroes(n int) int {
	r := 0
	for n >= 5 {
		n /= 5
		r += n
	}
	return r
}
