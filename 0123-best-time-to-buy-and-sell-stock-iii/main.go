package main

import "fmt"

func maxProfit2(prices []int) int {
	fmt.Println("prices", prices)

	memo := map[string]int{}

	var rec func(i, prev, left int) int
	rec = func(i, prev, left int) int {
		if i >= len(prices) {
			return 0
		}

		key := fmt.Sprintf("%d-%d-%d", i, prev, left)

		if _, exists := memo[key]; !exists {
			result := 0
			if prev == -1 {
				if left > 0 {
					// we can buy
					result = max(result, rec(i+1, prices[i], left-1))
				}
				// or skip
			} else {
				// we can sell
				result = max(result, prices[i]-prev+rec(i+1, -1, left))
			}
			// or skip
			result = max(result, rec(i+1, prev, left))
			memo[key] = result
		}

		return memo[key]
	}

	return rec(0, -1, 2)
}

func maxProfit(prices []int) int {
	// fmt.Println("Prices", prices)
	n := len(prices)

	mins := make([]int, n)
	leftProfit := make([]int, n)
	mins[0] = prices[0]
	for i := 1; i < n; i++ {
		mins[i] = min(mins[i-1], prices[i])
		leftProfit[i] = max(leftProfit[i-1], prices[i]-mins[i-1])
	}
	// fmt.Println("Mins", mins)
	// fmt.Println("Left profit", leftProfit)

	maxs := make([]int, n)
	rightProfit := make([]int, n)
	maxs[n-1] = prices[n-1]
	for i := n - 2; i >= 0; i-- {
		maxs[i] = max(prices[i], maxs[i+1])
		rightProfit[i] = max(rightProfit[i+1], maxs[i+1]-prices[i])
	}
	// fmt.Println("Maxs", maxs)
	// fmt.Println("Right profit", rightProfit)

	result := 0
	for i := 0; i < n-1; i++ {
		result = max(result, leftProfit[i]+rightProfit[i+1])
	}
	result = max(result, leftProfit[n-1])

	return result
}

func main() {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	fmt.Println("Example 1", maxProfit(prices)) // 6

	prices = []int{1, 2, 3, 4, 5}
	fmt.Println("Example 2", maxProfit(prices)) // 4

	prices = []int{7, 6, 4, 3, 1}
	fmt.Println("Example 3", maxProfit(prices)) // 0
}
