package Num_17_letterCombinations

var numToByteMap = map[byte][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	if len(digits) == 1 {
		return numToByteMap[digits[0]]
	}
	var result []string
	r := letterCombinations(digits[:len(digits)-1])
	for _, v := range numToByteMap[digits[len(digits)-1]] {
		for i := 0; i < len(r); i++ {
			result = append(result, r[i]+v)
		}
	}
	return result
}
