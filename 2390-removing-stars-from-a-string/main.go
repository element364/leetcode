package main

import "fmt"

func removeStars(s string) string {
	stack := []rune{}
	for _, c := range s {
		if c == '*' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}
	return string(stack)
}

func main() {
	fmt.Println("Example 1", removeStars("leet**cod*e")) // "lecoe"
	fmt.Println("Example 2", removeStars("erase*****"))  // ""
}
