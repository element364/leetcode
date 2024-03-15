package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	items    []int
	minValue int
}

func Constructor() MinStack {
	return MinStack{
		minValue: math.MaxInt,
	}
}

func (this *MinStack) Push(val int) {
	this.items = append(this.items, val)
	this.minValue = min(this.minValue, val)
}

func (this *MinStack) Pop() {
	top := this.Top()
	this.items = this.items[:len(this.items)-1]

	if top == this.minValue && len(this.items) > 0 {
		this.minValue = this.items[0]
		for i := 1; i < len(this.items); i++ {
			this.minValue = min(this.minValue, this.items[i])
		}
	}

	if len(this.items) == 0 {
		this.minValue = math.MaxInt
	}
}

func (this *MinStack) Top() int {
	return this.items[len(this.items)-1]
}

func (this *MinStack) GetMin() int {
	return this.minValue
}

func main() {
	fmt.Println("go")
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin()) // return -3
	minStack.Pop()
	fmt.Println(minStack.Top())    // return 0
	fmt.Println(minStack.GetMin()) // return -2
}
