package main

import "fmt"

func singleNumber2(nums []int) int {
	ones := 0
	twos := 0

	for _, num := range nums {
		ones ^= num & ^twos
		twos ^= num & ^ones
	}

	return ones
}

func singleNumber(nums []int) int {
	size := 33
	bits := make([]int, size)
	for _, num := range nums {
		fmt.Println("Num", num)
		bin := ""
		position := 0
		if num < 0 {
			bits[size-1]++
			num = -num
		}
		for num > 0 {
			bin += fmt.Sprintf("%d", num&1)
			bits[position] += num & 1
			num = num >> 1
			position++
		}
		fmt.Println("Num bin", bin)
	}
	fmt.Println("Bits before", bits)

	result := 0
	for i := 0; i < size; i++ {
		bits[i] %= 3
		if i == size-1 {
			if bits[i] != 0 {
				result = -result
			}
		} else {
			result = result + bits[i]<<i
		}
	}
	fmt.Println("Bits after", bits)

	return result
}

func main() {
	nums := []int{2, 2, 3, 2}
	fmt.Println(singleNumber(nums))

	nums = []int{0, 1, 0, 1, 0, 1, 99}
	fmt.Println(singleNumber(nums))
	return

	nums = []int{30000, 500, 100, 30000, 100, 30000, 100}
	fmt.Println(singleNumber(nums))

	nums = []int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}
	fmt.Println(singleNumber(nums))

	nums = []int{43, 16, 45, 89, 45, -2147483648, 45, 2147483646, -2147483647, -2147483648, 43, 2147483647, -2147483646, -2147483648, 89, -2147483646, 89, -2147483646, -2147483647, 2147483646, -2147483647, 16, 16, 2147483646, 43}
	fmt.Println(singleNumber(nums))

	nums = []int{-401451, -177656, -2147483646, -473874, -814645, -2147483646, -852036, -457533, -401451, -473874, -401451, -216555, -917279, -457533, -852036, -457533, -177656, -2147483646, -177656, -917279, -473874, -852036, -917279, -216555, -814645, 2147483645, -2147483648, 2147483645, -814645, 2147483645, -216555}
	fmt.Println(singleNumber(nums))
}
