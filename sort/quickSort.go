package sort

func QuickSort(s []int) []int {
	//界限，如果数组长度小于等于2 返回排序后的数组
	if len(s) <= 1 {
		return s
	}
	key := s[0]
	references := []int{key}
	left, right := make([]int, 0), make([]int, 0)
	for i := 1; i < len(s); i++ {
		switch true {
		case s[i] > key:
			right = append(right, s[i])
			continue
		case s[i] == key:
			references = append(references, s[i])
			continue
		default:
			left = append(left, s[i])
		}
	}
	result := append(QuickSort(left), references...)
	result = append(result, QuickSort(right)...)
	return result
}
