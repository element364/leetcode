package main

import "fmt"

func singleNumber(nums []int) int {
	result := 0

	for _, n := range nums {
		result ^= n
	}

	return result
}

func main() {
	nums := []int{2, 2, 1}
	fmt.Println("Example 1", singleNumber(nums))

	nums = []int{4, 1, 2, 1, 2}
	fmt.Println("Example 2", singleNumber(nums))

	nums = []int{1}
	fmt.Println("Example 3", singleNumber(nums))
}
