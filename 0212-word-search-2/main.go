package main

import "fmt"

type Trie struct {
	children [26]*Trie
	wordEnd  bool
}

func (this *Trie) Add(word string) {
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
	var dfs func(node *Trie, current string)
	dfs = func(node *Trie, current string) {
		if node.wordEnd {
			fmt.Println(current)
		}
		for i := range this.children {
			if node.children[i] != nil {
				dfs(node.children[i], current+string('a'+i))
			}
		}
	}
	dfs(this, "")
}

func findWords(board [][]byte, words []string) []string {
	trie := &Trie{}
	for _, word := range words {
		trie.Add(word)
	}

	n := len(board)
	m := len(board[0])
	used := make([][]bool, n)
	for i := 0; i < n; i++ {
		used[i] = make([]bool, m)
	}

	foundMap := map[string]bool{}

	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}

	var dfs func(x, y int, node *Trie, current string)
	dfs = func(x, y int, node *Trie, current string) {
		if node.wordEnd {
			foundMap[current] = true
		}

		for i := range dx {
			nx := x + dx[i]
			ny := y + dy[i]
			if nx >= 0 && nx < n && ny >= 0 && ny < m && !used[nx][ny] {
				ni := board[nx][ny] - 'a'

				if node.children[ni] != nil {
					used[x][y] = true
					dfs(nx, ny, node.children[ni], current+string('a'+ni))
					used[x][y] = false
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			index := board[i][j] - 'a'
			if trie.children[index] != nil {
				used[i][j] = true
				dfs(i, j, trie.children[index], string(board[i][j]))
				used[i][j] = false
			}
		}
	}

	result := []string{}
	for word := range foundMap {
		result = append(result, word)
	}
	return result
}

func main() {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}
	fmt.Println(findWords(board, words))

	board = [][]byte{
		{'a', 'b'},
		{'c', 'd'},
	}
	words = []string{"abcb"}
	fmt.Println(findWords(board, words))

	board = [][]byte{
		{'o', 'a', 'b', 'n'},
		{'o', 't', 'a', 'e'},
		{'a', 'h', 'k', 'r'},
		{'a', 'f', 'l', 'v'},
	}
	words = []string{"oa", "oaa"}
	fmt.Println(findWords(board, words))
}
