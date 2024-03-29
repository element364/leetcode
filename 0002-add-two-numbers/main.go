package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (this *ListNode) Print() {
	fmt.Print("[")
	curr := this
	for curr != nil {
		fmt.Print(curr.Val)
		if curr.Next != nil {
			fmt.Print(", ")
		}
		curr = curr.Next
	}
	fmt.Println("]")
}

func sliceToList(nums []int) *ListNode {
	var head *ListNode = nil
	var curr *ListNode = nil
	for _, num := range nums {
		node := &ListNode{
			Val: num,
		}
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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode = nil
	var curr *ListNode = nil

	c1 := l1
	c2 := l2
	carry := 0

	for c1 != nil || c2 != nil || carry > 0 {
		sum := carry
		if c1 != nil {
			sum += c1.Val
			c1 = c1.Next
		}
		if c2 != nil {
			sum += c2.Val
			c2 = c2.Next
		}

		carry = sum / 10
		node := &ListNode{
			Val: sum % 10,
		}
		if result == nil {
			result = node
			curr = result
		} else {
			curr.Next = node
			curr = node
		}
	}

	if result == nil {
		return &ListNode{
			Val: 0,
		}
	}
	return result
}

func main() {
	fmt.Println("sliceToList")
	l := []int{1, 2, 3, 4, 5}
	sliceToList(l).Print()

	fmt.Println("Example 1")
	l1 := []int{2, 4, 3}
	l2 := []int{5, 6, 4}
	addTwoNumbers(sliceToList(l1), sliceToList(l2)).Print() // [7, 0, 8]

	fmt.Println("Example 2")
	l1 = []int{0}
	l2 = []int{0}
	addTwoNumbers(sliceToList(l1), sliceToList(l2)).Print() // [0]

	fmt.Println("Example 3")
	l1 = []int{9, 9, 9, 9, 9, 9, 9}
	l2 = []int{9, 9, 9, 9}
	addTwoNumbers(sliceToList(l1), sliceToList(l2)).Print() // [8,9,9,9,0,0,0,1]

	fmt.Println("Example 4")
	l1 = []int{1, 6, 0, 3, 3, 6, 7, 2, 0, 1}
	l2 = []int{6, 3, 0, 8, 9, 6, 6, 9, 6, 1}
	addTwoNumbers(sliceToList(l1), sliceToList(l2)).Print() // [7,9,0,1,3,3,4,2,7,2]
}
