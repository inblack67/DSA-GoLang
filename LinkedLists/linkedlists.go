package linkedlists

import (
	"fmt"
)

// Node ...
type Node struct {
	data int
	next *Node
}

// LinkedList ...
type LinkedList struct {
	head *Node
	size int
}

// Append ...
func Append(ll *LinkedList, data int) {

	newNode := Node{
		data: data,
		next: nil,
	}

	if ll.size == 0 {
		ll.head = &newNode
		ll.size++
		return
	}

	currentNode := ll.head
	for i := 0; i < ll.size-1; i++ {
		currentNode = currentNode.next
	}

	currentNode.next = &newNode
	ll.size++
}

// Prepend ...
func Prepend(ll *LinkedList, data int) {
	newNode := Node{
		data: data,
	}
	newNode.next = ll.head
	ll.head = &newNode
	ll.size++
}

// Pop ...
func Pop(ll *LinkedList) {
	if ll.size == 0 {
		return
	}
	if ll.size == 1 {
		ll.head = nil
		return
	}
	currentNode := ll.head
	var prevNode *Node = nil
	for i := 0; i < ll.size-1; i++ {
		prevNode = currentNode
		currentNode = currentNode.next
	}
	prevNode.next = nil
	ll.size--
}

// Createlinkedlist ...
func Createlinkedlist() *LinkedList {
	myList := new(LinkedList)
	myList.head = nil
	myList.size = 0
	return myList
}

// DisplayLinkedList ...
func DisplayLinkedList(ll LinkedList) {
	currentNode := ll.head
	for i := 0; i < ll.size; i++ {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}
