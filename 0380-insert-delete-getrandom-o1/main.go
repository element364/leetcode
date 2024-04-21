package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	set map[int]struct{}
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		set: make(map[int]struct{}),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, exists := this.set[val]; !exists {
		this.set[val] = struct{}{}
		return true
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, exists := this.set[val]; exists {
		delete(this.set, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	randNum := rand.Intn(len(this.set))
	counter := 0
	for k := range this.set {
		if counter == randNum {
			return k
		}
		counter++
	}
	return -1
}

func main() {
	fmt.Println("Example 1")
	randomizedSet := Constructor()
	fmt.Println(randomizedSet.Insert(1))   // Inserts 1 to the set. Returns true as 1 was inserted successfully.
	fmt.Println(randomizedSet.Remove(2))   // Returns false as 2 does not exist in the set.
	fmt.Println(randomizedSet.Insert(2))   // Inserts 2 to the set, returns true. Set now contains [1,2].
	fmt.Println(randomizedSet.GetRandom()) // getRandom() should return either 1 or 2 randomly.
	fmt.Println(randomizedSet.Remove(1))   // Removes 1 from the set, returns true. Set now contains [2].
	fmt.Println(randomizedSet.Insert(2))   // 2 was already in the set, so return false.
	fmt.Println(randomizedSet.GetRandom()) // Since 2 is the only number in the set, getRandom() will always return 2.
}
