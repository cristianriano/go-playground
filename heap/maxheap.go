package heap

import (
	"fmt"
)

type MaxHeap struct {
	Values []int
	size int
}

func NewMaxHeap() MaxHeap {
	return MaxHeap{}
}

func (heap *MaxHeap) Push(n int) {
	heap.Values = append(heap.Values, n)
	heap.size++

	for i := heap.size - 1; i >= 0; i = (i - 1)/2 {
		if heap.Values[i] > heap.Values[(i-1)/2] {
			heap.swap(i, (i-1)/2)
		} else {
			break
		}
	}
}

func (heap *MaxHeap) Pop() int {
	if heap.size == 0 { return -1 }

	val := heap.Values[0]
	heap.swap(0, heap.size - 1)
	heap.size--
	heapify(heap, 0)

	return val
}

func HeapSort(values []int) {
	heap := HeapifyArray(values)
	heap.Print()
	l := len(values) - 1
	for i := 0; i <= l; i++ {
		swap(values, 0, l - i)
		heap.size--
		heapify(heap, 0)
	}
}

func HeapifyArray(values []int) *MaxHeap {
	heap := &(MaxHeap{Values: values, size: len(values)})

	for i := heap.size - 1; i >= 0; i-- {
		heapify(heap, i)
	}

	return heap
}

func swap(values []int, i, j int) {
	values[i], values[j] = values[j], values[i]
}

func heapify(heap *MaxHeap, index int) {
	leftIndex, rightIndex := heap.leftChild(index), heap.rightChild(index)

	if leftIndex < 0 && rightIndex < 0 { return }
	left, right, val := heap.get(leftIndex), heap.get(rightIndex), heap.get(index)

	if val >= left && val >= right {
		return
	} else if left > right {
		heap.swap(leftIndex, index)
		heapify(heap, leftIndex)
	} else {
		heap.swap(rightIndex, index)
		heapify(heap, rightIndex)
	}
}

func (heap *MaxHeap) Size() int {
	return heap.size
}

func (heap *MaxHeap) Print() {
	fmt.Println(heap.Values[:heap.size])
}

func (heap *MaxHeap) get(i int) int {
	if i < 0 || i >= heap.size {
		return -1
	}
	return heap.Values[i]
}

func (heap *MaxHeap) swap(i, j int) {
	heap.Values[i], heap.Values[j] = heap.Values[j], heap.Values[i]
}

func (heap *MaxHeap) leftChild(i int) int {
	child := (2 * i) + 1
	if child >= heap.size {
		return -1
	}
	return child
}

func (heap *MaxHeap) rightChild(i int) int {
	child := (2 * i) + 2
	if child >= heap.size {
		return -1
	}
	return child
}

func (heap *MaxHeap) parent(i int) int {
	return (i - 1)/2
}