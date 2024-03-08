package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	indices := map[int]int{}
	for i, num := range nums {
		if ind, exists := indices[num]; exists {
			if i-ind <= k {
				return true
			}
		}
		indices[num] = i

	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	k := 3
	fmt.Println(containsNearbyDuplicate(nums, k)) // true

	nums = []int{1, 0, 1, 1}
	k = 1
	fmt.Println(containsNearbyDuplicate(nums, k)) // true

	nums = []int{1, 2, 3, 1, 2, 3}
	k = 2
	fmt.Println(containsNearbyDuplicate(nums, k)) // false
}
