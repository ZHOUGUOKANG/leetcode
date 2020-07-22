package Num_13_romanToInt

var NumeralRomanDict = map[string]int{
	"M":  1000,
	"CM": 900,
	"D":  500,
	"CD": 400,
	"C":  100,
	"XC": 90,
	"L":  50,
	"XL": 40,
	"X":  10,
	"IX": 9,
	"V":  5,
	"IV": 4,
	"I":  1,
}

func RomanToInt(s string) int {
	num := 0
	length := len(s)
	tmpStr := ""
	for i := 0; i < length; i++ {
		tmpStr += string(s[i])
		if i < length-1 {
			if _, ok := NumeralRomanDict[tmpStr+string(s[i+1])]; ok {
				i++
				tmpStr += string(s[i])
			}
		}
		num += NumeralRomanDict[tmpStr]
		tmpStr = ""
	}
	return num
}
