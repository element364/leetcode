package main

import "fmt"

func romanToInt(s string) int {
	result := 0

	m := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := 0; i < len(s); i++ {
		c := rune(s[i])
		if i+1 < len(s) && m[c] < m[rune(s[i+1])] {
			result -= m[c]
		} else {
			result += m[c]
		}
	}
	return result
}

func main() {
	s := "III"
	fmt.Println(romanToInt(s)) // 3

	s = "LVIII"
	fmt.Println(romanToInt(s)) // 58

	s = "MCMXCIV"
	fmt.Println(romanToInt(s)) // 1994
}
