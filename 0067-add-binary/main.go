package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	maxLen := max(len(a), len(b)) + 1

	result := ""
	var carry byte = 0
	i := len(a) - 1
	j := len(b) - 1
	for k := 0; k < maxLen; k++ {
		var ab byte = 0
		if i >= 0 {
			ab = a[i] - '0'
		}
		var bb byte = 0
		if j >= 0 {
			bb = b[j] - '0'
		}

		sum := ab + bb + carry
		if sum > 1 {
			sum = sum % 2
			carry = 1
		} else {
			carry = 0
		}

		if k == maxLen-1 {
			if sum == 1 {
				result = strconv.Itoa(int(sum)) + result
			}
		} else {
			result = strconv.Itoa(int(sum)) + result
		}

		i--
		j--
	}

	return result
}

func main() {
	a := "11"
	b := "1"
	fmt.Println(addBinary(a, b))

	a = "1010"
	b = "1011"
	fmt.Println(addBinary(a, b))

	a = "0"
	b = "0"
	fmt.Println(addBinary(a, b))

	a = "1"
	b = "0"
	fmt.Println(addBinary(a, b))

	a = "0"
	b = "1"
	fmt.Println(addBinary(a, b))

	a = "1"
	b = "1"
	fmt.Println(addBinary(a, b))
}
