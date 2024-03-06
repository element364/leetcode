package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	letters := map[rune]int{}
	for _, c := range magazine {
		letters[c]++
	}

	for _, c := range ransomNote {
		letters[c]--
		if letters[c] < 0 {
			return false
		}
	}

	return true
}

func main() {
	ransomNote := "a"
	magazine := "b"
	fmt.Println(canConstruct(ransomNote, magazine))

	ransomNote = "aa"
	magazine = "ab"
	fmt.Println(canConstruct(ransomNote, magazine))

	ransomNote = "aa"
	magazine = "aab"
	fmt.Println(canConstruct(ransomNote, magazine))
}
