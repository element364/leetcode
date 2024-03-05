package main

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {
	sMap := map[byte]byte{}
	tMap := map[byte]byte{}
	for i := 0; i < len(s); i++ {
		if sMap[s[i]] != tMap[t[i]] {
			return false
		}
		sMap[s[i]] = byte(i + 1)
		tMap[t[i]] = byte(i + 1)
	}
	return true
}

func main() {
	s := "egg"
	t := "add"
	fmt.Println(isIsomorphic(s, t))

	s = "foo"
	t = "bar"
	fmt.Println(isIsomorphic(s, t))

	s = "paper"
	t = "title"
	fmt.Println(isIsomorphic(s, t))

	s = "bbbaaaba"
	t = "aaabbbba"
	fmt.Println(isIsomorphic(s, t))

	t = "badc"
	t = "baba"
	fmt.Println(isIsomorphic(s, t))
}
