package DynamicProgramming

import "testing"

func TestBackPack(t *testing.T) {
	weight := map[string]int{"water": 3, "book": 1, "food": 2, "jack": 2, "camera": 1}
	value := map[string]int{"water": 10, "book": 3, "food": 9, "jack": 5, "camera": 6}
	output := backpack(weight, value, 6)
	println(output)
}
