package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	visited := map[*ListNode]bool{}
	current := head
	for current != nil {
		visited[current] = true
		if visited[current.Next] {
			return true
		}
		current = current.Next
	}

	return false
}

func main() {
	fmt.Println("Kek")
	ln14 := ListNode{
		Val: -4,
	}
	ln13 := ListNode{
		Val: 0,
	}
	ln12 := ListNode{
		Val: 2,
	}
	ln11 := ListNode{
		Val: 3,
	}
	ln11.Next = &ln12
	ln12.Next = &ln13
	ln13.Next = &ln14
	ln14.Next = &ln12
	fmt.Println(hasCycle(&ln11))

	ln21 := ListNode{
		Val: 1,
	}
	ln22 := ListNode{
		Val: 2,
	}
	ln21.Next = &ln22
	ln22.Next = &ln21
	fmt.Println(hasCycle(&ln21))

	ln31 := ListNode{
		Val: 1,
	}
	fmt.Println(hasCycle(&ln31))
}
