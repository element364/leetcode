package main

import "fmt"

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	j := 0
	for i := 0; i < len(t); i++ {
		// fmt.Println("i", string(t[i]))
		if s[j] == t[i] {
			// fmt.Println("Matched", string(t[i]), "at", i, "with", string(s[j]), "at", j)
			j++
		}
		if j >= len(s) {
			return true
		}
	}
	return false
}

func main() {
	s := "abc"
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))

	s = "axc"
	t = "ahbgdc"
	fmt.Println(isSubsequence(s, t))

	s = ""
	t = "ahbgdc"
	fmt.Println(isSubsequence(s, t))

	s = "ahbgdc"
	t = ""
	fmt.Println(isSubsequence(s, t))
}
