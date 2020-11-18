package Num_134_CanCompleteCircuit

func canCompleteCircuit(gas []int, cost []int) int {
	for i, n := 0, len(gas); i < n; {
		sum, cnt := 0, 0
		for cnt < n {
			j := (i + cnt) % n
			sum += gas[j] - cost[j]
			if sum < 0 {
				break
			}
			cnt++
		}
		if cnt == n {
			return i
		} else {
			i += cnt + 1
		}
	}
	return -1
}
