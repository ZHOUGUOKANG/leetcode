package DynamicProgramming

import (
	"fmt"
)

func backpack(weight map[string]int, value map[string]int, capacity int) int {
	celi := make([][]int, 0)
	maxValue := 0
	for lastGoods, w := range weight {
		score := make([]int, capacity)
		//i 是背包容量
		for i := 1; i <= capacity; i++ {
			if weight[lastGoods] <= i {
				score[i-1] = value[lastGoods]
				remainderCapacity := i - weight[lastGoods]
				if remainderCapacity > 0 {
					maxScore := 0
					for _, l := range celi {
						if l[remainderCapacity-1] > maxScore {
							maxScore = l[remainderCapacity-1]
						}
					}
					score[i-1] += maxScore
				}
			} else {
				maxScore := 0
				if i-2 >= 0 {
					maxScore = score[i-2]
					for _, l := range celi {
						if l[i-2] > maxScore {
							maxScore = l[i-2]
						}
					}
				}
				score[i-1] = maxScore
			}
			if score[i-1] > maxValue {
				maxValue = score[i-1]
			}
		}
		fmt.Printf("%10s %d %d %v\n", lastGoods, w, value[lastGoods], score)
		celi = append(celi, score)
	}
	return maxValue
}
