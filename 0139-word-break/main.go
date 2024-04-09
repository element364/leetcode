package main

import (
	"fmt"
	"strings"
)

type Trie struct {
	children [26]*Trie
	wordEnd  bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	current := this
	for _, c := range word {
		index := c - 'a'
		if current.children[index] == nil {
			current.children[index] = &Trie{}
		}
		current = current.children[index]
	}
	current.wordEnd = true
}

func (this *Trie) Print() {
	fmt.Println("Trie")
	var dfs func(node *Trie, s string)
	dfs = func(node *Trie, s string) {
		if node.wordEnd {
			fmt.Println(" ", s)
		}
		for i, c := range node.children {
			if c != nil {
				dfs(c, s+string('a'+i))
			}
		}
	}
	dfs(this, "")
}

// func wordBreak(s string, wordDict []string) bool {
// 	trie := Constructor()
// 	for _, word := range wordDict {
// 		trie.Insert(word)
// 	}

// 	result := false
// 	found := make([]bool, len(s))

// 	var dfs func(p int, node *Trie)
// 	dfs = func(p int, node *Trie) {
// 		if result {
// 			return
// 		}
// 		if p == len(s) {
// 			if node.wordEnd {
// 				result = true
// 			}
// 			return
// 		}
// 		if node.wordEnd {
// 			found[p-1] = true

// 			left := s[p:]
// 			if found[len(left)-1] && strings.HasPrefix(s, left) {
// 				result = true
// 				return
// 			}
// 		}

// 		index := s[p] - 'a'
// 		if node.children[index] != nil {
// 			dfs(p+1, node.children[index])
// 			if node.children[index].wordEnd {
// 				dfs(p+1, &trie)
// 			}
// 		}
// 	}
// 	dfs(0, &trie)

// 	return result
// }

func wordBreak(s string, wordDict []string) bool {
	queue := []string{s}
	memo := map[string]bool{}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if _, exists := memo[current]; exists {
			continue
		}

		for _, word := range wordDict {
			if word == current {
				return true
			}
			if strings.HasPrefix(current, word) {
				queue = append(queue, current[len(word):])
				memo[current] = true
			}
		}
	}
	return false
}

func main() {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	fmt.Println("Example 1", wordBreak(s, wordDict)) // true

	s = "applepenapple"
	wordDict = []string{"apple", "pen"}
	fmt.Println("Example 2", wordBreak(s, wordDict)) // true

	s = "catsandog"
	wordDict = []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println("Example 3", wordBreak(s, wordDict)) // false

	s = "aaaaaaa"
	wordDict = []string{"aaaa", "aa"}
	fmt.Println("Example 4", wordBreak(s, wordDict)) // false

	s = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	wordDict = []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	fmt.Println("Example 5", wordBreak(s, wordDict)) // false
}
