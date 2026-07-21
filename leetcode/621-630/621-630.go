package _21_630

import (
	"container/heap"
	"sort"
)

// leetcode problem No. 622
type MyCircularQueue struct {
	buffer []int
	head   int
	rear   int
	k      int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		buffer: make([]int, k+1), // one redundant space
		head:   0,
		rear:   0, // rear point to place to insert value
		k:      k, // the actual useful size
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.buffer[this.rear] = value
	this.rear = (this.rear + 1) % (this.k + 1)
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = (this.head + 1) % (this.k + 1)
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.buffer[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.buffer[this.k+this.rear%(this.k+1)]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.head == this.rear
}

func (this *MyCircularQueue) IsFull() bool {
	return (this.rear+1)%(this.k+1) == this.head
}

// leetcode problem No. 630

type MaxHeap []int

func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j] // since it's max heap, use >
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return x
}

func scheduleCourse(courses [][]int) int {
	endingDay := 0
	maxHeap := &MaxHeap{}
	// sort by deadline ascending, we choose deadline because deadline is fixed while start time is not
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	for _, course := range courses {
		heap.Push(maxHeap, course[0])
		endingDay += course[0] // if taking this course, what the ending day will be

		if endingDay > course[1] { // if the ending day extend the deadline
			endingDay -= heap.Pop(maxHeap).(int) // not choose the course with max duration
		}
	}
	return maxHeap.Len()
}
