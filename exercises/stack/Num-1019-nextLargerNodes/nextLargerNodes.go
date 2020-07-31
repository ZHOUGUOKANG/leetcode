package Num_1019_nextLargerNodes

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {
	data := make([]int, 0)
	for head != nil {
		data = append(data, head.Val)
		head = head.Next
	}

	stack, ans := make([]int, len(data)), make([]int, len(data))

	for i := 0; i < len(data); i++ {
		if len(stack) == 0 {
			stack = append(stack, i)
		} else {
			for len(stack) != 0 && data[i] > data[stack[len(stack)-1]] {
				pop := len(stack) - 1
				ans[stack[pop]] = data[i]
				stack = stack[:pop]
			}
			stack = append(stack, i)
		}
	}
	return ans
}
