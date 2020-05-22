package sort

/**
归并排序
归并排序是建立在归并操作上的一种有效的排序算法。
该算法是采用分治法（Divide and Conquer）的一个非常典型的应用。
将已有序的子序列合并，得到完全有序的序列；即先使每个子序列有序，再使子序列段间有序。
若将两个有序表合并成一个有序表，称为2-路归并。
*/
func MergeSort(s []int) []int {
	length := len(s)
	if length < 2 {
		return s
	}
	mid := length / 2
	left := s[:mid]
	right := s[mid:]
	return merge(MergeSort(left), MergeSort(right))
}
func merge(left, right []int) []int {
	leftLen := len(left)
	rightLen := len(right)
	if leftLen == 0 {
		return right
	}
	if rightLen == 0 {
		return left
	}
	i, j := 0, 0
	output := make([]int, 0, leftLen+rightLen)
	for i < leftLen || j < rightLen {
		println(i, j, left[i], right[j])
		if left[i] < right[j] {
			output = append(output, left[i])
			i++
			if i >= leftLen {
				output = append(output, right[j:]...)
				break
			}
		} else {
			output = append(output, right[j])
			j++
			if j >= rightLen {
				output = append(output, left[i:]...)
				break
			}
		}

	}
	return output
}
