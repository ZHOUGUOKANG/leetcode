package sort

/**
堆排序
一般都用数组来表示堆，i结点的父结点下标就为(i – 1) / 2。它的左右子结点下标分别为2 * i + 1和2 * i + 2。如第0个结点左右子结点下标分别为1和2。
*/
func HeapSort(s []int) []int {
	length := len(s)
	if length < 2 {
		return s
	}
	s = sort(s)
	for i := 0; i < length; i++ {
		s = append(s[:i], sort(s[i:])...)
	}
	return s
}

func push(s []int, ins int) []int {
	s = append(s, ins)
	length := len(s)
	j := length - 1
	for {
		i := (j - 1) / 2 // parent
		if i == j {
			break
		}
		if s[i] > s[j] {
			s[i], s[j] = s[j], s[i]
			j = i
		} else {
			break
		}
	}
	return s
}

func sort(s []int) []int {
	length := len(s)
	output := make([]int, 0, length)
	for i := 0; i < length; i++ {
		output = push(output, s[i])
	}
	return output
}
