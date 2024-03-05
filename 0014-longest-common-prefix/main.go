package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	result := ""
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[0][i] != strs[j][i] {
				return result
			}
		}
		result = result + string(strs[0][i])
	}
	return result
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
	strs = []string{"dog", "racecar", "car"}

	fmt.Println(longestCommonPrefix(strs))
	strs = []string{"ab", "a"}
	fmt.Println(longestCommonPrefix(strs))
}
