package main

import "fmt"

func letterCombinations(digits string) []string {
	n := len(digits)

	mapping := map[rune][]rune{
		'2': []rune{'a', 'b', 'c'},
		'3': []rune{'d', 'e', 'f'},
		'4': []rune{'g', 'h', 'i'},
		'5': []rune{'j', 'k', 'l'},
		'6': []rune{'m', 'n', 'o'},
		'7': []rune{'p', 'q', 'r', 's'},
		'8': []rune{'t', 'u', 'v'},
		'9': []rune{'w', 'x', 'y', 'z'},
	}

	combinations := []string{}
	current := ""

	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			if len(current) > 0 {
				combinations = append(combinations, current)
			}
			return
		}
		c := rune(digits[i])
		for _, n := range mapping[c] {
			current += string(n)
			dfs(i + 1)
			current = current[:len(current)-1]
		}
	}

	dfs(0)

	return combinations
}

func main() {
	digits := "23"
	fmt.Println("Example 1", letterCombinations(digits)) // ["ad","ae","af","bd","be","bf","cd","ce","cf"]
	digits = ""
	fmt.Println("Example 2", letterCombinations(digits)) // []
	digits = "2"
	fmt.Println("Example 3", letterCombinations(digits)) // ["a","b","c"]
}
