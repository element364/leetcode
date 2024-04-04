package main

import (
	"fmt"
)

func myPow(x float64, n int) float64 {
	var result float64 = 1
	absn := n
	if absn < 0 {
		absn = -absn
	}
	for absn > 0 {
		if absn&1 == 1 {
			result *= x
			absn--
		} else {
			x *= x
			absn >>= 1
		}
	}

	if n < 0 {
		return 1 / result
	}
	return result
}

func main() {
	x := 2.00000
	n := 10
	fmt.Println("Example 1", myPow(x, n))

	x = 2.10000
	n = 3
	fmt.Println("Example 2", myPow(x, n))

	x = 2.00000
	n = -2
	fmt.Println("Example 3", myPow(x, n))
}
