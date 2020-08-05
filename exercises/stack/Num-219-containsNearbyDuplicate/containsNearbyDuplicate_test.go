package Num_219_containsNearbyDuplicate

import "testing"

func TestContainsNearbyDuplicate(t *testing.T) {
	nums, k := []int{1, 2, 3, 1}, 3
	out := containsNearbyDuplicate(nums, k)
	if true != out {
		t.Error(nums, k, true, out)
	}
}

func TestContainsNearbyDuplicate1(t *testing.T) {
	nums, k := []int{1, 0, 1, 1}, 1
	out := containsNearbyDuplicate(nums, k)
	if true != out {
		t.Error(nums, k, true, out)
	}
}
func TestContainsNearbyDuplicate2(t *testing.T) {
	nums, k := []int{1, 2, 3, 1, 2, 3}, 2
	out := containsNearbyDuplicate(nums, k)
	if false != out {
		t.Error(nums, k, false, out)
	}
}

func TestContainsNearbyDuplicate3(t *testing.T) {
	nums, k := []int{99, 99}, 2
	out := containsNearbyDuplicate(nums, k)
	if true != out {
		t.Error(nums, k, true, out)
	}
}
func TestContainsNearbyDuplicate4(t *testing.T) {
	nums, k := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9}, 3
	out := containsNearbyDuplicate(nums, k)
	if true != out {
		t.Error(nums, k, true, out)
	}
}
