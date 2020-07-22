package Num_12_IntToRoman

//罗马数字对应关系
var RomanNumeralDict = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}
var RomanSort = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 5, 4, 1}

func IntToRoman(num int) string {
	output := ""
	for _, value := range RomanSort {
		merchant := num / value
		if merchant > 0 {
			for i := 0; i < merchant; i++ {
				output += RomanNumeralDict[value]
			}
			num = num - merchant*value
		}
	}
	return output
}
