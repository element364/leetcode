package main

import "fmt"

func maxArea(height []int) int {
	n := len(height)

	result := 0
	i := 0
	j := n - 1
	for i < j {
		h := min(height[i], height[j])
		w := j - i
		s := h * w
		result = max(result, s)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return result
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println("Example 1", maxArea(height))

	height = []int{1, 1}
	fmt.Println("Example 2", maxArea(height))
}
