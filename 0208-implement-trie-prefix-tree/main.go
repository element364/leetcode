package main

import "fmt"

type Trie struct {
	children [26]*Trie
	wordEnd  bool
}

func Constructor() Trie {
	return Trie{}
}

func getIndex(c rune) rune {
	return c - 'a'
}

func (this *Trie) Insert(word string) {
	// fmt.Println("Inserting", word)
	current := this
	for _, c := range word {
		index := getIndex(c)
		if current.children[index] == nil {
			// fmt.Println("  Creating node for", string(c))
			current.children[index] = &Trie{}
		}
		current = current.children[index]
	}
	// fmt.Println("  Setting wordEnd for", current)
	current.wordEnd = true
}

func (this *Trie) Search(word string) bool {
	current := this
	for _, c := range word {
		index := getIndex(c)
		if current.children[index] == nil {
			return false
		}
		current = current.children[index]
		// fmt.Println("  Moving to letter", string(c), current.wordEnd)
	}
	return current.wordEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	current := this
	for _, c := range prefix {
		index := getIndex(c)
		if current.children[index] == nil {
			return false
		}
		current = current.children[index]
	}
	return true
}

func main() {
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple")) // return True
	fmt.Println(trie.Search("app"))   // return False
	trie.Insert("az")
	fmt.Println(trie.StartsWith("app")) // return True
	trie.Insert("app")
	fmt.Println(trie.Search("app")) // return True
}
