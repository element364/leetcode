package main

import (
	"fmt"
	"strconv"
)

func toInt(s string) int {
	result, _ := strconv.Atoi(s)
	return result
}

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			o1 := stack[len(stack)-1]
			o2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, o2+o1)
			case "-":
				stack = append(stack, o2-o1)
			case "*":
				stack = append(stack, o2*o1)
			case "/":
				stack = append(stack, o2/o1)
			}
		} else {
			stack = append(stack, toInt(token))
		}
	}
	return stack[0]
}

func main() {
	tokens := []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(tokens))

	tokens = []string{"4", "13", "5", "/", "+"}
	fmt.Println(evalRPN(tokens))

	tokens = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(tokens))
}
