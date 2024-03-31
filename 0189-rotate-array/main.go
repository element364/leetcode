package main

import "fmt"

func rotate(nums []int, k int) {
	k = k % len(nums)
	for i := 0; i < k; i++ {
		prev := nums[0]
		for j := 1; j < len(nums); j++ {
			tmp := nums[j]
			nums[j] = prev
			prev = tmp
		}
		nums[0] = prev
	}
}

func main() {
	fmt.Println("Example 1")
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	fmt.Println(nums)

	fmt.Println("Example 2")
	nums = []int{-1, -100, 3, 99}
	k = 2
	rotate(nums, k)
	fmt.Println(nums)
}
