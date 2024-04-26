package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	var result *ListNode
	var prev *ListNode

	var adding bool
	var candidate *int

	for current := head; current != nil; current = current.Next {
		if candidate == nil {
			adding = true
			candidate = &current.Val
		} else {
			if current.Val == *candidate {
				adding = false
			} else {
				if adding {
					next := &ListNode{
						Val: *candidate,
					}
					if result == nil {
						result = next
						prev = result
					} else {
						prev.Next = next
						prev = next
					}
				}
				adding = true
				candidate = &current.Val
			}
		}
	}
	if adding {
		next := &ListNode{
			Val: *candidate,
		}
		if result == nil {
			result = next
			prev = result
		} else {
			prev.Next = next
			prev = next
		}
	}

	return result
}

func sliceToList(s []int) *ListNode {
	var head *ListNode
	var prev *ListNode
	for _, n := range s {
		if prev == nil {
			head = &ListNode{
				Val: n,
			}
			prev = head
		} else {
			next := &ListNode{
				Val: n,
			}
			prev.Next = next
			prev = next
		}
	}
	return head
}

func printList(head *ListNode) {
	fmt.Print("List [")
	curr := head
	for curr != nil {
		fmt.Print(curr.Val, ",")
		curr = curr.Next
	}
	fmt.Println("]")
}

func main() {
	fmt.Println("Example 0")
	head := sliceToList([]int{})
	printList(head)
	head = deleteDuplicates(head)
	printList(head)

	fmt.Println("Example 1")
	head = sliceToList([]int{1, 2, 3, 3, 4, 4, 5})
	printList(head)
	head = deleteDuplicates(head)
	printList(head)

	fmt.Println("Example 2")
	head = sliceToList([]int{1, 1, 1, 2, 3})
	printList(head)
	head = deleteDuplicates(head)
	printList(head)
}
