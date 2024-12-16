package priorityqueue

import (
	"container/heap"
)

// Item represents an element in the priority queue.
type Item[T any] struct {
	Value    T
	Priority int
	Index    int // The index of the item in the heap.
}

// PriorityQueue implements a generic priority queue.
type PriorityQueue[T any] struct {
	elements []*Item[T]
}

// New creates a new PriorityQueue.
func New[T any]() *PriorityQueue[T] {
	return &PriorityQueue[T]{elements: []*Item[T]{}}
}

// Len returns the number of elements in the queue.
func (pq PriorityQueue[T]) Len() int {
	return len(pq.elements)
}

// Less compares two elements based on their priority.
func (pq PriorityQueue[T]) Less(i, j int) bool {
	// Higher priority items come first.
	return pq.elements[i].Priority < pq.elements[j].Priority
}

// Swap swaps two elements in the queue.
func (pq PriorityQueue[T]) Swap(i, j int) {
	pq.elements[i], pq.elements[j] = pq.elements[j], pq.elements[i]
	pq.elements[i].Index = i
	pq.elements[j].Index = j
}

// Push adds an element to the queue.
func (pq *PriorityQueue[T]) Push(x any) {
	item := x.(*Item[T])
	item.Index = len(pq.elements)
	pq.elements = append(pq.elements, item)
}

// Pop removes and returns the highest-priority element from the queue.
func (pq *PriorityQueue[T]) Pop() any {
	old := pq.elements
	n := len(old)
	item := old[n-1]
	item.Index = -1 // Mark as removed.
	pq.elements = old[0 : n-1]
	return item
}

// Update modifies the priority and value of an element in the queue.
func (pq *PriorityQueue[T]) Update(item *Item[T], value T, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}

// PushItem adds an item to the priority queue.
func (pq *PriorityQueue[T]) PushItem(value T, priority int) {
	item := &Item[T]{Value: value, Priority: priority}
	heap.Push(pq, item)
}

// PopItem removes and returns the highest-priority item from the queue.
func (pq *PriorityQueue[T]) PopItem() T {
	item := heap.Pop(pq).(*Item[T])
	return item.Value
}

// Peek returns the highest-priority item without removing it from the queue.
func (pq *PriorityQueue[T]) Peek() *Item[T] {
	if len(pq.elements) == 0 {
		return nil
	}
	return pq.elements[0]
}
