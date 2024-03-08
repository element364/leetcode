package main

import (
	"fmt"
	"reflect"
)

func stringToCounts(s string) map[rune]int {
	result := map[rune]int{}

	for _, c := range s {
		result[c]++
	}

	return result
}

func isAnagram(s string, t string) bool {
	sc := stringToCounts(s)
	tc := stringToCounts(t)
	return reflect.DeepEqual(sc, tc)
}

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))

	s = "rat"
	t = "car"
	fmt.Println(isAnagram(s, t))
}
