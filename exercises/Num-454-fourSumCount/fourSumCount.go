package Num_454_FourSumCount

func fourSumCount(A []int, B []int, C []int, D []int) (ans int) {
	countAB := map[int]int{}
	for _, v := range A {
		for _, w := range B {
			countAB[v+w]++
		}
	}
	for _, v := range C {
		for _, w := range D {
			ans += countAB[-v-w]
		}
	}
	return
}
