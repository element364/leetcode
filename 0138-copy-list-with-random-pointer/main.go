package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	var result *Node
	var pointer *Node

	randomMap := map[*Node]*Node{}

	current := head
	for current != nil {
		cp := &Node{Val: current.Val}
		randomMap[current] = cp
		if result == nil {
			result = cp
			pointer = result
		} else {
			pointer.Next = cp
			pointer = cp
		}
		current = current.Next
	}

	// Set random
	p1 := head
	p2 := result
	for p1 != nil {
		p2.Random = randomMap[p1.Random]
		p1 = p1.Next
		p2 = p2.Next
	}

	return result
}

func (this *Node) Print() {
	current := this
	for current != nil {
		item := []int{}
		item = append(item, current.Val)
		if current.Random != nil {
			item = append(item, current.Random.Val)
		}
		fmt.Println(item)
		current = current.Next
	}
}

func main() {
	fmt.Println("Example 1")
	n1 := &Node{Val: 7}
	n2 := &Node{Val: 13}
	n3 := &Node{Val: 11}
	n4 := &Node{Val: 10}
	n5 := &Node{Val: 1}
	n1.Next = n2
	n2.Next = n3
	n2.Random = n1
	n3.Next = n4
	n3.Random = n5
	n4.Next = n5
	n4.Random = n3
	n5.Random = n1
	n1.Print()
	fmt.Println("Copy")
	cp := copyRandomList(n1)
	cp.Print()

	fmt.Println("Example 2")
	n1 = &Node{Val: 1}
	n2 = &Node{Val: 2}
	n1.Next = n2
	n1.Random = n2
	n2.Random = n2
	n1.Print()
	fmt.Println("Copy")
	cp = copyRandomList(n1)
	cp.Print()

	fmt.Println("Example 3")
	n1 = &Node{Val: 3}
	n2 = &Node{Val: 3}
	n3 = &Node{Val: 3}
	n1.Next = n2
	n2.Next = n3
	n2.Random = n1
	n1.Print()
	fmt.Println("Copy")
	cp = copyRandomList(n1)
	cp.Print()
}
