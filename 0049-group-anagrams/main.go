package main

import (
	"fmt"
	"strconv"
)

func strToCountsStr(s *string) string {
	counts := [26]int{}
	for _, c := range *s {
		code := c - 'a'
		counts[code]++
	}

	result := ""
	for i := range counts {
		if counts[i] > 0 {
			result += string('a'+i) + strconv.Itoa(counts[i])
		}
	}
	return result
}

func groupAnagrams(strs []string) [][]string {
	anagrams := map[string][]string{}

	for _, s := range strs {
		hash := strToCountsStr(&s)
		anagrams[hash] = append(anagrams[hash], s)
	}

	result := make([][]string, 0, len(anagrams))
	for _, v := range anagrams {
		result = append(result, v)
	}
	return result
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs)) // [["bat"],["nat","tan"],["ate","eat","tea"]]

	strs = []string{""}
	fmt.Println(groupAnagrams(strs)) // [[""]]

	strs = []string{"a"}
	fmt.Println(groupAnagrams(strs)) // [[""]]
}
