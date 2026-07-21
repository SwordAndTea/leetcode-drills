package _371_380

import (
	"container/heap"
)

// leetcode problem No. 374
func topKFrequent(nums []int, k int) []int {
	numCounter := make(map[int]int)
	maxFreq := 0
	for _, num := range nums {
		numCounter[num]++
		if numCounter[num] > maxFreq {
			maxFreq = numCounter[num]
		}
	}

	buckets := make([][]int, maxFreq+1)
	for num, count := range numCounter {
		buckets[count] = append(buckets[count], num)
	}

	result := make([]int, 0, k)
	for i := maxFreq; i >= 1; i-- {
		bucket := buckets[i]
		if len(bucket) != 0 {
			if len(result)+len(bucket) < k {
				result = append(result, bucket...)
			} else {
				result = append(result, bucket[:k-len(result)]...)
			}
		}
	}
	return result
}

// leetcode problem No. 373
type ValuePair struct {
	sum    int
	indexI int
	indexJ int
}

type MinHeap []*ValuePair

func (h *MinHeap) Len() int { return len(*h) }

func (h *MinHeap) Less(i, j int) bool {
	return (*h)[i].sum <= (*h)[j].sum
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*ValuePair))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	v := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return v
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	ans := make([][]int, k)
	minHeap := &MinHeap{}

	heap.Init(minHeap)
	heap.Push(minHeap, &ValuePair{sum: nums1[0] + nums2[0], indexI: 0, indexJ: 0})
	visited := make(map[int]map[int]bool)
	visited[0] = make(map[int]bool)
	visited[0][0] = true
	for kk := 0; kk < k; kk++ {
		curPair := heap.Pop(minHeap).(*ValuePair)
		ans[kk] = []int{nums1[curPair.indexI], nums2[curPair.indexJ]}
		if nextI := curPair.indexI + 1; nextI < len(nums1) {
			if visited[nextI] == nil {
				visited[nextI] = make(map[int]bool)
			}
			if !visited[nextI][curPair.indexJ] {
				heap.Push(minHeap, &ValuePair{
					sum:    nums1[nextI] + nums2[curPair.indexJ],
					indexI: nextI,
					indexJ: curPair.indexJ,
				})
				visited[nextI][curPair.indexJ] = true
			}
		}
		if nextJ := curPair.indexJ + 1; nextJ < len(nums2) {
			if visited[curPair.indexI] == nil {
				visited[curPair.indexI] = make(map[int]bool)
			}
			if !visited[curPair.indexI][nextJ] {
				heap.Push(minHeap, &ValuePair{
					sum:    nums1[curPair.indexI] + nums2[nextJ],
					indexI: curPair.indexI,
					indexJ: nextJ,
				})
				visited[curPair.indexI][nextJ] = true
			}
		}
	}
	return ans
}
