package heaps

import "fmt"

type MaxHeap struct {
	arr []int
}

func GetParentIndex(index int) int {
	return (index - 1) / 2
}

// Left child will always be the odd index => right => even
func GetLeftChild(index int) int {
	return (2 * index) + 1
}

func GetRightChild(index int) int {
	return (2 * index) + 2
}

func (h *MaxHeap) Swap(index1, index2 int) {
	h.arr[index1], h.arr[index2] = h.arr[index2], h.arr[index1]
}

func (h *MaxHeap) Heapify(index int) {
	for h.arr[GetParentIndex((index))] < h.arr[index] {
		h.Swap(GetParentIndex(index), index)
		// index = GetParentIndex(index)
	}
}

func (heap *MaxHeap) Insert(el int) {
	heap.arr = append(heap.arr, el)
	heap.Heapify(len(heap.arr) - 1)
}

func (heap *MaxHeap) HeapifyDown(index int) {
	lastIndex := len(heap.arr) - 1
	leftIndex, rightIndex := GetLeftChild(lastIndex), GetRightChild(lastIndex)
	childToCompare := 0
	for leftIndex <= lastIndex {
		if leftIndex == lastIndex {
			// no right child
			childToCompare = leftIndex
		} else if heap.arr[leftIndex] > heap.arr[rightIndex] {
			childToCompare = leftIndex
		} else {
			childToCompare = rightIndex
		}
		if heap.arr[index] < heap.arr[childToCompare] {
			heap.Swap(index, childToCompare)
			index = childToCompare
			leftIndex, rightIndex = GetLeftChild(lastIndex), GetRightChild(lastIndex)
		} else {
			return
		}
	}
}

// Get Max Element => and remove it from the heap
func (heap *MaxHeap) Extract() (int, error) {
	if len(heap.arr) == 0 {
		return -1, fmt.Errorf("heap is empty")
	}
	root := heap.arr[0]
	lastIndex := len(heap.arr) - 1
	heap.arr[0] = heap.arr[lastIndex]
	heap.arr = heap.arr[:lastIndex] // removing last element
	heap.HeapifyDown(0)
	return root, nil
}
