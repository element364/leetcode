package main

import "fmt"

func maxProfit(prices []int) int {
	maxProfit := 0
	if len(prices) < 2 {
		return maxProfit
	}

	nextMax := prices[len(prices)-1]
	for i := len(prices) - 2; i >= 0; i-- {
		maxProfit = max(maxProfit, nextMax-prices[i])
		nextMax = max(nextMax, prices[i])
	}
	return maxProfit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))

	prices = []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))

	prices = []int{}
	fmt.Println(maxProfit(prices))

	prices = []int{1}
	fmt.Println(maxProfit(prices))
}
