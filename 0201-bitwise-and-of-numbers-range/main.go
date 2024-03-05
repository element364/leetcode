package main

import (
	"fmt"
	"slices"
)

func toBits(n int) []uint8 {
	bits := []uint8{}
	for n > 0 {
		bits = append(bits, uint8(n&1))
		n = n >> 1
	}

	slices.Reverse(bits)

	return bits
}

func toInt(bits []uint8) int {
	result := 0
	for i := 0; i < len(bits); i++ {
		result = result<<1 + int(bits[i])
	}
	return result
}

func rangeBitwiseAnd(left int, right int) int {
	leftB := toBits(left)
	rightB := toBits(right)

	for len(leftB) < len(rightB) {
		leftB = append([]uint8{0}, leftB...)
	}

	// fmt.Println("Left", left, "\n", leftB)
	// fmt.Println("Right", right, "\n", rightB)

	resultB := make([]uint8, len(rightB))
	for i := len(leftB) - 1; i >= 0; i-- {
		if leftB[i] == 1 && rightB[i] == 1 {
			noOnesLeft := true
			for k := 0; k < i; k++ {
				if leftB[k] != rightB[k] {
					noOnesLeft = false
					break
				}
			}

			if noOnesLeft {
				resultB[i] = 1
			}
		}
	}
	// fmt.Println("Result", toInt(resultB), "\n", resultB)
	return toInt(resultB)
}

func main() {
	fmt.Println(rangeBitwiseAnd(5, 7))              // 4
	fmt.Println(rangeBitwiseAnd(0, 0))              // 0
	fmt.Println(rangeBitwiseAnd(1, 2147483647))     // 0
	fmt.Println(rangeBitwiseAnd(1, 2))              // 0
	fmt.Println(rangeBitwiseAnd(20000, 2147483647)) // 0
	fmt.Println(rangeBitwiseAnd(3, 6))              // 0
}
