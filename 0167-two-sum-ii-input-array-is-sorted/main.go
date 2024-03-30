package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers)-1; i++ {
		rest := target - numbers[i]
		l := i + 1
		r := len(numbers) - 1

		for r-l > 1 {
			m := (l + r) / 2
			if numbers[m] < rest {
				l = m
			} else {
				r = m
			}
		}
		if numbers[l] == rest {
			return []int{i + 1, l + 1}
		}
		if numbers[r] == rest {
			return []int{i + 1, r + 1}
		}
	}
	return []int{}
}

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	fmt.Println("Example 1", twoSum(numbers, target))

	numbers = []int{2, 3, 4}
	target = 6
	fmt.Println("Example 2", twoSum(numbers, target))

	numbers = []int{-1, 0}
	target = -1
	fmt.Println("Example 1", twoSum(numbers, target))
}
