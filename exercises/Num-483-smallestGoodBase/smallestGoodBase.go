package Num_483_SmallestGoodBase

import (
	"math"
	"math/bits"
	"strconv"
)

func smallestGoodBase(n string) string {
	intN, _ := strconv.Atoi(n)
	fN := float64(intN)
	for m := bits.Len(uint(intN)) - 1; m > 1; m-- {
		k := int(math.Pow(fN, 1/float64(m)))
		mul, sum := 1, 1
		for i := 0; i < m; i++ {
			mul *= k
			sum += mul
		}
		if sum == intN {
			return strconv.Itoa(k)
		}
	}
	return strconv.Itoa(intN - 1)
}
