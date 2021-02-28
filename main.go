package main

import (
	linkedlists "github.com/inblack67/dsa/LinkedLists"
)

func main() {
	ll := linkedlists.Createlinkedlist()
	linkedlists.Append(ll, 1)
	linkedlists.Append(ll, 2)
	linkedlists.Append(ll, 3)
	linkedlists.DisplayLinkedList(*ll)
}
