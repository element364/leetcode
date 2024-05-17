package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	s := fmt.Sprintf("%d", x)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func main() {
	x := 121
	fmt.Println("Example 1", isPalindrome(x)) // true

	x = -121
	fmt.Println("Example 2", isPalindrome(x)) // false

	x = 10
	fmt.Println("Example 3", isPalindrome(x)) // false
}
