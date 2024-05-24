package main

import (
	"fmt"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func (this *Node) Print() {
	queue := []*Node{}
	if this != nil {
		queue = append(queue, this)
	}
	for len(queue) > 0 {
		level := []*Node{}
		for currentLen := len(queue); currentLen > 0; currentLen-- {
			current := queue[0]
			queue = queue[1:]
			level = append(level, current)
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
		fmt.Print("[")
		for _, n := range level {
			if n.Next != nil {
				fmt.Printf("%d->%d ,", n.Val, n.Next.Val)
			} else {
				fmt.Printf("%d, ", n.Val)
			}
		}
		fmt.Println("]")
	}
}

func connect(root *Node) *Node {
	queue := []*Node{}
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		level := []*Node{}
		for currentLen := len(queue); currentLen > 0; currentLen-- {
			current := queue[0]
			queue = queue[1:]
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
			level = append(level, current)
		}
		for i := 0; i < len(level)-1; i++ {
			level[i].Next = level[i+1]
		}
	}
	return root
}

func main() {
	fmt.Println("Example 1")
	root := &Node{
		Val: 1,
		Left: &Node{
			Val:   2,
			Left:  &Node{Val: 4},
			Right: &Node{Val: 5},
		},
		Right: &Node{
			Val:   3,
			Right: &Node{Val: 7},
		},
	}
	root.Print()
	root = connect(root)
	root.Print()

	fmt.Println("Example 2")
	root = nil
	root.Print()
	root = connect(root)
	root.Print()
}
