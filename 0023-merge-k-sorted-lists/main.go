package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sliceToList(s []int) *ListNode {
	var head *ListNode
	var curr *ListNode

	for _, n := range s {
		node := &ListNode{Val: n}
		if head == nil {
			head = node
			curr = head
		} else {
			curr.Next = node
			curr = node
		}
	}

	return head
}

func (this *ListNode) Print() {
	fmt.Print("[")
	for curr := this; curr != nil; curr = curr.Next {
		fmt.Print(curr.Val, ",")
	}
	fmt.Println("]")
}

func mergeKLists(lists []*ListNode) *ListNode {
	var head *ListNode
	var curr *ListNode

	for {
		iMin := -1
		for i := 0; i < len(lists); i++ {
			if lists[i] != nil {
				if iMin == -1 || lists[i].Val < lists[iMin].Val {
					iMin = i
				}
			}
		}

		if iMin == -1 {
			break
		}

		node := &ListNode{Val: lists[iMin].Val}
		lists[iMin] = lists[iMin].Next

		if head == nil {
			head = node
			curr = head
		} else {
			curr.Next = node
			curr = node
		}
	}

	return head
}

func main() {
	fmt.Println("Example 1")
	lists := []*ListNode{
		sliceToList([]int{1, 4, 5}),
		sliceToList([]int{1, 3, 4}),
		sliceToList([]int{2, 6}),
	}
	merged := mergeKLists(lists)
	merged.Print()

	fmt.Println("Example 2")
	lists = []*ListNode{}
	merged = mergeKLists(lists)
	merged.Print()

	fmt.Println("Example 3")
	lists = []*ListNode{
		sliceToList([]int{}),
	}
	merged = mergeKLists(lists)
	merged.Print()
}
