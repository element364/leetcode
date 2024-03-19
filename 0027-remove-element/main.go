package main

import "fmt"

func removeElement(nums []int, val int) int {
	head := 0
	tail := len(nums) - 1
	for head <= tail {
		if nums[head] == val {
			nums[head], nums[tail] = nums[tail], nums[head]
			tail--
		} else {
			head++
		}
	}
	return head
}

func main() {
	fmt.Println("Example 1")
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println(removeElement(nums, val))
	fmt.Println(nums)

	fmt.Println("Example 2")
	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	val = 2
	fmt.Println(removeElement(nums, val))
	fmt.Println(nums)

	fmt.Println("Example 3")
	nums = []int{4, 4, 4, 4, 4}
	val = 4
	fmt.Println(removeElement(nums, val))
	fmt.Println(nums)

	fmt.Println("Example 4")
	nums = []int{}
	val = 1
	fmt.Println(removeElement(nums, val))
	fmt.Println(nums)
}
