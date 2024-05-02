package main

import "fmt"

func binSearch(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for r-l > 1 {
		m := (l + r) / 2
		if nums[m] > target {
			r = m
		} else {
			l = m
		}
	}

	if nums[l] == target {
		return l
	}
	if nums[r] == target {
		return r
	}
	return -1
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	m := binSearch(nums, target)
	fmt.Println("Found m", m)
	if m == -1 {
		return []int{-1, -1}
	}

	l := m
	for l >= 0 && nums[l] == target {
		l--
	}

	r := m
	for r < len(nums) && nums[r] == target {
		r++
	}
	return []int{l + 1, r - 1}
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println("Example 1", searchRange(nums, target)) // 3 4

	nums = []int{5, 7, 7, 8, 8, 10}
	target = 6
	fmt.Println("Example 2", searchRange(nums, target)) // -1 -1

	nums = []int{}
	target = 0
	fmt.Println("Example 3", searchRange(nums, target)) // -1 -1

	nums = []int{5}
	target = 5
	fmt.Println("Example 5", searchRange(nums, target)) // 0 0

	nums = []int{6}
	target = 5
	fmt.Println("Example 6", searchRange(nums, target)) // -1 -1

	nums = []int{7, 7, 7, 7, 7, 7, 7}
	target = 7
	fmt.Println("Example 7", searchRange(nums, target)) // 0 6

	nums = []int{1, 2, 3}
	target = 1
	fmt.Println("Example 8", searchRange(nums, target)) // 0, 0

	nums = []int{0, 0, 1, 2, 2}
	target = 2
	fmt.Println("Example 9", searchRange(nums, target)) // [3,4]
}
