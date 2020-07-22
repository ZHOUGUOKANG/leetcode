package Num_offer_11_minArray

//func minArray(numbers []int) int {
//	sort.Ints(numbers)
//	return numbers[0]
//}
func minArray(numbers []int) int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		mid := left + (right-left)/2
		if numbers[mid] > numbers[right] {
			left = mid + 1
		} else if numbers[mid] < numbers[right] {
			right = mid
		} else {
			right--
		}
	}
	return numbers[left]
}
