package Num_27_removeElement

func removeElement(nums []int, val int) int {
	n := len(nums)
	i, j := 0, n-1
	for ; i <= j; i++ {
		if nums[i] == val {
			flag := true
			for ; j > i; j-- {
				if nums[j] != val {
					nums[i], nums[j] = nums[j], val
					flag = false
					break
				}
			}
			if flag {
				return j
			}
		}
	}
	return i
}
