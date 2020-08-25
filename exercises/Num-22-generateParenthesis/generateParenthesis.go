package Num_22_generateParenthesis

import "fmt"

//var cache map[int][]string
//func generateParenthesis(n int) []string {
//	if cache == nil{
//		cache = make(map[int][]string)
//	}
//	return generate(n)
//}
//func generate(n int)[]string  {
//	if v,ok := cache[n];ok{
//		return v
//	}
//	var ans []string
//	if n == 0{
//		ans = append(ans,"")
//	}else {
//		for c:=0;c<n;c++{
//			a := generate(c)
//			b := generate(n - 1 - c)
//			for _,x := range a{
//				for _,y := range b{
//					ans = append(ans,"(" + x + ")" + y)
//				}
//			}
//		}
//	}
//	cache[n] = ans
//	return ans
//}
var cache map[int][]string

func generateParenthesis(n int) []string {
	if cache == nil {
		cache = make(map[int][]string)
	}
	return gen(n)
}
func gen(n int) []string {
	if v, ok := cache[n]; ok {
		return v
	}
	var ans []string
	if n == 0 {
		ans = append(ans, "")
	} else {
		for i := 0; i < n; i++ {
			before := gen(i)
			after := gen(n - i - 1)
			for _, x := range before {
				for _, y := range after {
					ans = append(ans, fmt.Sprintf("(%s)%s", x, y))
				}
			}
		}
	}
	cache[n] = ans
	return ans
}
