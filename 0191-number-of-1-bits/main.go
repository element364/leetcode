package main

import "fmt"

func hammingWeight(num uint32) int {
	var result uint8 = 0
	for num > 0 {
		result += uint8(num & 1)
		num >>= 1
	}
	return int(result)
}

func main() {
	fmt.Println(hammingWeight(11))
}
