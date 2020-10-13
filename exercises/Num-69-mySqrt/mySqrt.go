package Num_69_MySqrt

//func mySqrt(x int) int {
//	i := 1
//	for{
//		if i*i > x{
//			return i-1
//		}
//		i++
//	}
//}
func mySqrt(x int) int {
	l, r := 0, x
	mid, num := 0, 0
	for r >= l {
		mid = (l + r) / 2
		num = mid * mid
		if num > x {
			r = mid - 1
		} else if num < x {
			l = mid + 1
		} else {
			return mid
		}
	}
	return r
}
