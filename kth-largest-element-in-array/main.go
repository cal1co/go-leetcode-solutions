package main

import (
	"container/heap"
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("1", reflect.DeepEqual(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4), 4))
	fmt.Println("2", reflect.DeepEqual(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2), 5))
}

type MinHeap []int

func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

// Find Kth Largest with MinHeap Implementation
//
// Time complexity: O(n * log(k))
func findKthLargest(nums []int, k int) int {
	h := MinHeap{}
	for _, val := range nums {
		heap.Push(&h, val)
		if len(h) > k {
			heap.Pop(&h)
		}
	}
	return heap.Pop(&h).(int)
}
