package main

import "fmt"

func longestConsecutive(nums []int) int {
	hash := map[int]int{}
	for _, num := range nums {
		hash[num] = 0
	}

	result := 0
	for k := range hash {
		count := 1
		current := k
		// fmt.Println("Counting from", current)
		for {
			next := current + 1
			if v, exists := hash[next]; exists {
				if v > 0 {
					count += v
					// fmt.Println("Early stop", count)
					break
				}
				count++
				current = next
			} else {
				break
			}
		}
		hash[k] = count
		result = max(result, count)
	}

	return result
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums)) // 4

	nums = []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(longestConsecutive(nums)) // 9
}
