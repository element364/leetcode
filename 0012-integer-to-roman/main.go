package main

import "fmt"

// 1000 - M
// 900 - CM
// 500 - D
// 400 - CD
// 100 - C
// 90 - XC
// 50 - L
// 40 - XL
// 10 - X
// 9 - IX
// 5 - V
// 4 - IV
// 1 - I

func intToRoman(num int) string {
	result := ""

	keys := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	for i := range keys {
		// fmt.Println("KV", keys[i], values[i])
		for num >= values[i] {
			result += keys[i]
			num -= values[i]
			// fmt.Println("Taking", values[i], result)
		}
	}

	return result
}

func main() {
	num := 3
	fmt.Println("Example 1", intToRoman(num)) // III

	num = 58
	fmt.Println("Example 2", intToRoman(num)) // LVIII

	num = 1994
	fmt.Println("Example 2", intToRoman(num)) // MCMXCIV
}
