package _21_630

import (
	"container/heap"
	"sort"
)

// leetcode problem No. 630

type MaxHeap []int

func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
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
		endingDay += course[0]

		if endingDay > course[1] { // if the ending day extend the deadline
			endingDay -= heap.Pop(maxHeap).(int) // not choose the course with max duration
		}
	}
	return maxHeap.Len()
}
