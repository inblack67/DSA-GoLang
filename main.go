package main

import (
	"fmt"

	heaps "github.com/inblack67/dsa/Heaps"
)

func main() {
	// ll := linkedlists.Createlinkedlist()
	// linkedlists.Append(ll, 2)
	// linkedlists.Append(ll, 3)
	// linkedlists.Prepend(ll, 1)
	// linkedlists.Append(ll, 4)
	// linkedlists.Pop(ll)
	// linkedlists.DisplayLinkedList(*ll)

	m := &heaps.MaxHeap{}

	heapData := []int{10, 20, 30}
	for _, v := range heapData {
		m.Insert(v)
		fmt.Println(m)
	}
}
