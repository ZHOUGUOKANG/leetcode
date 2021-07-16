package Num_275_HIndex

func hIndex(citations []int) (ans int) {
	n := len(citations)
	for i := 0; i < n; i++ {
		if t := n - i; citations[i] >= t {
			return t
		}
	}
	return
}
