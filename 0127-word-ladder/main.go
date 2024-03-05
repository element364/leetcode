package main

import (
	"fmt"
	"slices"
)

type QueueItem struct {
	Index int
	Steps int
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	beginWordIndex := slices.Index(wordList, beginWord)
	if beginWordIndex == -1 {
		wordList = append([]string{beginWord}, wordList...)
		beginWordIndex = 0
	}

	reacheble := func(a, b int) bool {
		if len(wordList[a]) != len(wordList[b]) {
			return false
		}

		differs := 0
		for i := 0; i < len(wordList[a]); i++ {
			if wordList[a][i] != wordList[b][i] {
				differs++
			}
			if differs > 1 {
				return false
			}
		}
		return true
	}

	cache := make([][]bool, len(wordList))
	for i := 0; i < len(wordList); i++ {
		cache[i] = make([]bool, len(wordList))
	}
	for i := 0; i < len(wordList)-1; i++ {
		for j := i + 1; j < len(wordList); j++ {
			r := reacheble(i, j)
			cache[i][j] = r
			cache[j][i] = r
		}
	}

	queue := []QueueItem{{
		Index: beginWordIndex,
		Steps: 1,
	}}
	visited := map[string]bool{beginWord: true}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for i := 0; i < len(wordList); i++ {
			if !visited[wordList[i]] && cache[current.Index][i] {
				nextSteps := current.Steps + 1
				if wordList[i] == endWord {
					return nextSteps
				}
				visited[wordList[i]] = true
				queue = append(queue, QueueItem{
					Index: i,
					Steps: nextSteps,
				})
			}
		}
	}

	return 0
}

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Println(ladderLength(beginWord, endWord, wordList))

	beginWord = "hit"
	endWord = "cog"
	wordList = []string{"hot", "dot", "dog", "lot", "log"}
	fmt.Println(ladderLength(beginWord, endWord, wordList))

	beginWord = "hot"
	endWord = "dog"
	wordList = []string{"hot", "dog", "dot"}
	fmt.Println(ladderLength(beginWord, endWord, wordList))
}
