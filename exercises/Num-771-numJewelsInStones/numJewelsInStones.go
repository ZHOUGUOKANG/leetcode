package Num_771_NumJewelsInStones

func numJewelsInStones(J string, S string) int {
	js := make(map[byte]bool)
	for i := 0; i < len(J); i++ {
		js[J[i]] = true
	}
	count := 0
	for i := 0; i < len(S); i++ {
		if _, ok := js[S[i]]; ok {
			count++
		}
	}
	return count
}
func numJewelsInStones1(J string, S string) int {
	count := 0
	for i := 0; i < len(J); i++ {
		c := J[i]
		for j := 0; j < len(S); j++ {
			if c == S[j] {
				count++
			}
		}
	}
	return count
}
