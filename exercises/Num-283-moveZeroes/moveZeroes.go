package Num_283_MoveZeroes

func moveZeroes(nums []int) {
	slow, fast := 0, 0
	for ; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
	}
	for ; slow < len(nums); slow++ {
		nums[slow] = 0
	}
}
