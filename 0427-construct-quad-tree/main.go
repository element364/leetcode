package main

import "fmt"

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func (this *Node) Print() {
	var dfs func(node *Node, depth int)
	dfs = func(node *Node, depth int) {
		if node == nil {
			return
		}
		for i := 0; i < depth; i++ {
			fmt.Print(" ")
		}
		if node.Val {
			fmt.Print(1)
		} else {
			fmt.Print(0)
		}
		if node.IsLeaf {
			fmt.Println("L")
		} else {
			fmt.Println("N")
		}
		dfs(node.TopLeft, depth+2)
		dfs(node.TopRight, depth+2)
		dfs(node.BottomLeft, depth+2)
		dfs(node.BottomRight, depth+2)
	}
	dfs(this, 0)

	fmt.Println("Print")
}

func construct(grid [][]int) *Node {
	var dfs func(x, y, l int) *Node
	dfs = func(x, y, l int) *Node {
		if l == 0 {
			return nil
		}

		count := 0
		for i := 0; i < l; i++ {
			for j := 0; j < l; j++ {
				count += grid[x+i][y+j]
			}
		}
		if count == l*l {
			return &Node{
				Val:    true,
				IsLeaf: true,
			}
		}
		if count == 0 {
			return &Node{
				Val:    false,
				IsLeaf: true,
			}
		}

		w := l >> 1
		return &Node{
			Val:         true,
			IsLeaf:      false,
			TopLeft:     dfs(x, y, w),
			TopRight:    dfs(x, y+w, w),
			BottomLeft:  dfs(x+w, y, w),
			BottomRight: dfs(x+w, y+w, w),
		}
	}
	l := len(grid)
	return dfs(0, 0, l)
}

func main() {
	fmt.Println("Example 1")
	grid := [][]int{
		{0, 1},
		{1, 0},
	}
	qt := construct(grid)
	qt.Print()

	fmt.Println("Example 2")
	grid = [][]int{
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
	}
	qt = construct(grid)
	qt.Print()
}
