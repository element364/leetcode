package main

import "fmt"

func climbStairs(n int) int {
	memo := map[int]int{
		1: 1,
		2: 2,
	}
	var rec func(n int) int
	rec = func(n int) int {
		if _, exists := memo[n]; !exists {
			memo[n] = rec(n-1) + rec(n-2)
		}
		return memo[n]
	}
	return rec(n)
}

func main() {
	n := 1
	fmt.Println("Example 2", climbStairs(n))

	n = 2
	fmt.Println("Example 1", climbStairs(n))

	n = 3
	fmt.Println("Example 2", climbStairs(n))
}
