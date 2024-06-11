package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	n := len(candies)
	result := make([]bool, n)

	for i, count := range candies {
		check := true
		for j, other := range candies {
			if j != i && count+extraCandies < other {
				check = false
				break
			}
		}
		result[i] = check
	}

	return result
}

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	fmt.Println("Example 1", kidsWithCandies(candies, extraCandies)) // [true,true,true,false,true]

	candies = []int{4, 2, 1, 1, 2}
	extraCandies = 1
	fmt.Println("Example 2", kidsWithCandies(candies, extraCandies)) // [true,false,false,false,false]

	candies = []int{12, 1, 12}
	extraCandies = 10
	fmt.Println("Example 3", kidsWithCandies(candies, extraCandies)) // [true,false,true]
}
