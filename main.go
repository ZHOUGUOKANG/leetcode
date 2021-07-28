package main

//func main() {
//	//sum1,sum2 := loop(10),recursion(10)
//	//fmt.Printf("sum1=%d,sum2=%d",sum1,sum2)
//}

func loop(n int) (sum int) {
	for i := 0; i < n; i++ {
		sum += i
	}
	return
}

func helper(i, n int) int {
	if i >= n {
		return 0
	}
	return i + helper(i+1, n)
}
