package main

import (
	"fmt"

	"github.com/emirpasic/gods/stacks/linkedliststack"
)

func isValid(s string) bool {
	stack := linkedliststack.New()
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack.Push(c)
		} else {
			top, ok := stack.Pop()
			if !ok {
				return false
			}

			if c == ')' && top != '(' {
				return false
			}
			if c == ']' && top != '[' {
				return false
			}
			if c == '}' && top != '{' {
				return false
			}
		}
	}
	return stack.Empty()
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
}
