package main

import "fmt"

type WordDictionary struct {
	children [26]*WordDictionary
	wordEnd  bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	current := this
	for _, c := range word {
		index := c - 'a'
		if current.children[index] == nil {
			current.children[index] = &WordDictionary{}
		}
		current = current.children[index]
	}
	current.wordEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	if len(word) == 0 {
		return this.wordEnd
	}

	if word[0] == '.' {
		for i := range this.children {
			if this.children[i] != nil {
				if this.children[i].Search(word[1:]) {
					return true
				}
			}
		}
	} else {
		index := word[0] - 'a'
		if this.children[index] != nil && this.children[index].Search(word[1:]) {
			return true
		}
	}

	return false
}

func (this *WordDictionary) Print() {
	results := []string{}

	var dfs func(root *WordDictionary, current string)
	dfs = func(root *WordDictionary, current string) {
		if root.wordEnd {
			results = append(results, current)
		}

		for i := range root.children {
			if root.children[i] != nil {
				dfs(root.children[i], current+string('a'+i))
			}
		}
	}
	dfs(this, "")

	fmt.Println("Print", results)
}

func main() {
	fmt.Println("Test 1")
	wordDictionary := Constructor()
	wordDictionary.AddWord("bad")
	wordDictionary.AddWord("dad")
	wordDictionary.AddWord("mad")
	wordDictionary.Print()
	fmt.Println(wordDictionary.Search("pad")) // return False
	fmt.Println(wordDictionary.Search("bad")) // return True
	fmt.Println(wordDictionary.Search(".ad")) // return True
	fmt.Println(wordDictionary.Search("b..")) // return True

	fmt.Println("Test 2")
	wd := Constructor()
	wd.AddWord("at")
	wd.AddWord("and")
	wd.AddWord("an")
	wd.AddWord("add")
	wd.Print()
	fmt.Println(wd.Search("a"))   // return false
	fmt.Println(wd.Search(".at")) // return false
	wd.AddWord("bat")
	wd.Print()
	fmt.Println(wd.Search(".at"))  // true
	fmt.Println(wd.Search("an."))  // true
	fmt.Println(wd.Search("a.d.")) // false
	fmt.Println(wd.Search("b."))   // false
	fmt.Println(wd.Search("a.d"))  // true
	fmt.Println(wd.Search("."))    // false
}
