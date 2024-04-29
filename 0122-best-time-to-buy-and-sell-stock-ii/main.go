package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	prc := []int{}
	for i, p := range prices {
		if i == 0 || p != prc[len(prc)-1] {
			prc = append(prc, p)
		}
	}

	n := len(prc)
	memo := make([]int, n)

	var dfs func(p, buyPrice int) int
	dfs = func(p, buyPrice int) int {
		if p >= n {
			// We finished
			return 0
		}

		if buyPrice == -1 && memo[p] > 0 {
			return memo[p]
		}

		best := -1
		if buyPrice == -1 {
			// We have no stocks, so we can buy for prices[i]
			buyingStock := -prc[p] + dfs(p+1, prc[p])
			if buyingStock > best {
				best = buyingStock
			}
		} else if prc[p] > buyPrice {
			// We have stock, so we can sell it now for prices[i]
			sellingStock := prc[p] + dfs(p+1, -1)
			if sellingStock > best {
				best = sellingStock
			}
		}

		skipping := dfs(p+1, buyPrice)
		if skipping > best {
			best = skipping
		}

		if buyPrice == -1 {
			memo[p] = best
		}

		return best
	}

	return dfs(0, -1)
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("Example 1", maxProfit(prices)) // 7

	prices = []int{1, 2, 3, 4, 5}
	fmt.Println("Example 2", maxProfit(prices)) // 4

	prices = []int{7, 6, 4, 3, 1}
	fmt.Println("Example 3", maxProfit(prices)) // 0
}
