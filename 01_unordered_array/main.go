package main

import "fmt"

type UnorderedArray struct {
	capacity     int
	lastPosition int
	values       []int
}

func NewUnorderedArray(capacity int) *UnorderedArray {
	return &UnorderedArray{
		capacity:     capacity,
		lastPosition: -1,
		values:       make([]int, capacity),
	}
}

func (ua *UnorderedArray) Print() {
	if ua.lastPosition == -1 {
		fmt.Println("The array is empty")
	} else {
		for i := 0; i <= ua.lastPosition; i++ {
			fmt.Println(i, " - ", ua.values[i])
		}
	}
}

func (ua *UnorderedArray) Insert(value int) {
	if ua.lastPosition == ua.capacity-1 {
		fmt.Println("Maximum capacity reached")
	} else {
		ua.lastPosition++
		ua.values[ua.lastPosition] = value
	}
}

func main() {
	ua := NewUnorderedArray(10)

	ua.Insert(3)
	ua.Insert(1)
	ua.Insert(4)
	ua.Insert(1)
	ua.Insert(5)

	ua.Print()
}
