package main

import (
	"fmt"
)

func coinChange(coins []int, amount int) int {
	visited := map[int]int{0: 0}
	queue := []int{0}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, coin := range coins {
			next := current + coin
			if next == amount {
				return visited[current] + 1
			}
			if next < amount {
				if _, ok := visited[next]; !ok {
					queue = append(queue, next)
					visited[next] = visited[current] + 1
				}
			}
		}
	}
	if v, ok := visited[amount]; ok {
		return v
	}
	return -1
}

func main() {
	// coins := []int{1, 2, 5}
	// amount := 11

	// coins := []int{2}
	// amount := 3

	// coins := []int{1}
	// amount := 0

	coins := []int{1, 2, 5}
	amount := 100

	result := coinChange(coins, amount)
	fmt.Println("result", result)
}
