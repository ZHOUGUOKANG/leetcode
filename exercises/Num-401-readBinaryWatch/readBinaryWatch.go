package Num_401_ReadBinaryWatch

//1,2,4,8 0~11
var hourMapping = map[int][]string{
	0: {"0"},
	1: {"1", "2", "4", "8"},
	2: {"3", "5", "9", "6", "10"},
	3: {"7", "11"},
}

//1,2,4,8,16,32 0~59
var minMapping = map[int][]string{
	0: {"00"},
	1: {"01", "02", "04", "08", "16", "32"},
	2: {"03", "05", "09", "17", "33", "06", "10", "18", "34", "12", "20", "36", "24", "40", "48"},
	3: {"07", "11", "19", "35", "14", "22", "38", "28", "44", "56", "13", "21", "37", "25", "41", "49", "26", "42", "50", "52"},
	4: {"15", "23", "39", "30", "46", "27", "43", "53", "29", "45", "46", "51", "54", "57", "58"},
	5: {"31", "47", "59", "55"},
}

func readBinaryWatch(turnedOn int) (ans []string) {
	if turnedOn >= 9 || turnedOn < 0 {
		return
	}
	cache := map[string]struct{}{}
	for i := 3; i >= 0; i-- {
		if j := turnedOn - i; j >= 0 {
			for _, h := range hourMapping[i] {
				for _, m := range minMapping[j] {
					cache[h+":"+m] = struct{}{}
				}
			}
		}
	}
	for k, _ := range cache {
		ans = append(ans, k)
	}
	return
}
