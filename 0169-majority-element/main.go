package main

import "fmt"

func majorityElement(nums []int) int {
	result := 0
	count := 0
	for _, num := range nums {
		if count == 0 {
			result = num
		}
		if num == result {
			count++
		} else {
			count--
		}
	}
	return result
}

func main() {
	nums := []int{3, 2, 3}
	fmt.Println("Example 1", majorityElement(nums))

	nums = []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println("Example 2", majorityElement(nums))
}
