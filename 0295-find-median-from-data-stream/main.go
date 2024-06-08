package main

import (
	"fmt"
	"sort"
)

type MedianFinder struct {
	s []int
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	// fmt.Println("Adding", num, "to", this.s)
	this.s = append(this.s, num)
}

func (this *MedianFinder) FindMedian() float64 {
	// fmt.Println("Finding median", this.s)
	sort.Ints(this.s)
	l := len(this.s)
	m := l >> 1
	// fmt.Println("m", m)
	if l&1 == 1 {
		return float64(this.s[m])
	} else {
		k := m - 1
		// fmt.Println("k", k)
		return float64(this.s[m]+this.s[k]) / 2
	}
}

func main() {
	fmt.Println("Example 1")
	medianFinder := Constructor()
	medianFinder.AddNum(1)                 // arr = [1]
	medianFinder.AddNum(2)                 // arr = [1, 2]
	fmt.Println(medianFinder.FindMedian()) // return 1.5 (i.e., (1 + 2) / 2)
	medianFinder.AddNum(3)                 // arr[1, 2, 3]
	fmt.Println(medianFinder.FindMedian()) // return 2.0
}
