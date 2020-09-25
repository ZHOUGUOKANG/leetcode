package Num_offer_08_01_waysToStep

func waysToStep(n int) int {
	switch n {
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 4
	default:
		a, b, c := 1, 2, 4
		var result int
		for i := 4; i <= n; i++ {
			result = (a + b + c) % 1000000007
			a, b, c = b, c, result
		}
		return result
	}
}
