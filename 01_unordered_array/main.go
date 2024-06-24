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

// O(1)
func (ua *UnorderedArray) Insert(value int) {
	if ua.lastPosition == ua.capacity-1 {
		fmt.Println("Maximum capacity reached")
	} else {
		ua.lastPosition++
		ua.values[ua.lastPosition] = value
	}
}

// O(n)
func (ua *UnorderedArray) Search(value int) int {
	for i := 0; i <= ua.lastPosition; i++ {
		if ua.values[i] == value {
			return i
		}
	}

	return -1
}

// O(n)
func (ua *UnorderedArray) Delete(value int) int {
	position := ua.Search(value)

	if position == -1 {
		return -1
	} else {
		for i := position; i < ua.lastPosition; i++ {
			ua.values[i] = ua.values[i+1]
		}

		ua.lastPosition--
		return position
	}
}

func main() {
	ua := NewUnorderedArray(10)

	fmt.Println("Insert")
	ua.Insert(3)
	ua.Insert(1)
	ua.Insert(4)
	ua.Insert(5)
	ua.Insert(8)

	ua.Print()

	fmt.Println("Search")
	fmt.Println("Index of 5:", ua.Search(5))
	fmt.Println("Index of 1:", ua.Search(1))
	fmt.Println("Index of 4:", ua.Search(4))

	fmt.Println("Search")
	ua.Delete(8)
	ua.Delete(4)

	ua.Print()

}
