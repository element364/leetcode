package main

import "fmt"

func calculate(s string) int {
	cur_stack := []int{}
	res_stack := []int{}

	sign := 1
	curr := 0
	result := 0
	for _, c := range s {
		switch c {
		case '+':
			result += curr * sign
			curr = 0
			sign = 1
			break
		case '-':
			result += curr * sign
			curr = 0
			sign = -1
			break
		case '(':
			cur_stack = append(cur_stack, result)
			res_stack = append(res_stack, sign)
			result = 0
			sign = 1
			break
		case ')':
			result += curr * sign
			curr = 0
			sign = res_stack[len(res_stack)-1]
			res_stack = res_stack[:len(res_stack)-1]

			result = sign*result + cur_stack[len(cur_stack)-1]
			cur_stack = cur_stack[:len(cur_stack)-1]
			break
		default:
			if c != ' ' {
				digit := int(c - '0')
				curr = curr*10 + digit
			}
		}
	}

	return result + curr*sign
}

func main() {
	s := "1 + 1"
	fmt.Println("Example 1", calculate(s)) // 2

	s = " 2-1 + 2 "
	fmt.Println("Example 2", calculate(s)) // 3

	s = "(1+(4+5+2)-3)+(6+8)"
	fmt.Println("Example 3", calculate(s)) // 23

	s = "1-(     -2)"
	fmt.Println("Example 4", calculate(s)) // 3
}
