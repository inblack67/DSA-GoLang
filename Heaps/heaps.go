package heaps

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
		index = GetParentIndex(index)
	}
}

func (heap *MaxHeap) Insert(el int) {
	heap.arr = append(heap.arr, el)
	heap.Heapify(len(heap.arr) - 1)
}
