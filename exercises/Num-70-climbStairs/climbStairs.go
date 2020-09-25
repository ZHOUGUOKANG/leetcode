package Num_70_ClimbStairs

func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	lst1, lst2 := 1, 2
	for i := 3; i <= n; i++ {
		lst1, lst2 = lst2, lst1+lst2
	}
	return lst2
}
