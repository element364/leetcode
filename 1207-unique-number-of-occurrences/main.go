package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
	m := map[int]int{}
	for _, item := range arr {
		m[item]++
	}

	o := map[int]bool{}
	for _, v := range m {
		if _, exists := o[v]; exists {
			return false
		}
		o[v] = true
	}
	return true
}

func main() {
	arr := []int{1, 2, 2, 1, 1, 3}
	fmt.Println("Example 1", uniqueOccurrences(arr))

	arr = []int{1, 2}
	fmt.Println("Example 2", uniqueOccurrences(arr))

	arr = []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}
	fmt.Println("Example 3", uniqueOccurrences(arr))
}
