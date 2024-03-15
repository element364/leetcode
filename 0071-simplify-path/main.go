package main

import (
	"fmt"
	"regexp"
	"strings"
)

func simplifyPath(path string) string {
	// fmt.Println("Simplify", path)
	stack := []string{}
	re := regexp.MustCompile("/+")

	for _, part := range re.Split(strings.TrimRight(path, "/"), -1) {
		// fmt.Println(" Part", part)
		if part == ".." {
			if len(stack) > 0 {
				stack = stack[0 : len(stack)-1]
			}
		} else if part != "." {
			stack = append(stack, part)
		}
		// fmt.Println("Current stack", stack)
	}

	result := strings.Join(stack, "/")
	if !strings.HasPrefix(result, "/") {
		return "/" + result
	}
	return result
}

func main() {
	path := "/home/"
	fmt.Println(simplifyPath(path))

	path = "/../"
	fmt.Println(simplifyPath(path))

	path = "/home//foo/"
	fmt.Println(simplifyPath(path))

	path = "/a/./b/../../c/"
	fmt.Println(simplifyPath(path))

	path = "/home/../../.."
	fmt.Println(simplifyPath(path))

	path = "/Users/element364/../Work/../leetcode/0071-simplify-path"
	fmt.Println(simplifyPath(path))

	path = "/Users/element364/../../Work/../leetcode/0071-simplify-path"
	fmt.Println(simplifyPath(path))
}
