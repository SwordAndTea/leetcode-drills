package _11_220

import (
	"container/heap"
	"sort"
)

type WordDictionary struct {
	IsWord   bool
	NextChar [26]*WordDictionary
}

func Constructor() WordDictionary {
	return WordDictionary{IsWord: false, NextChar: [26]*WordDictionary{}}
}

func (this *WordDictionary) AddWord(word string) {
	curLevel := this
	for _, w := range word {
		if curLevel.NextChar[w-'a'] == nil {
			curLevel.NextChar[w-'a'] = &WordDictionary{IsWord: false, NextChar: [26]*WordDictionary{}}
		}
		curLevel = curLevel.NextChar[w-'a']
	}
	curLevel.IsWord = true
}

func doSearch(curDictionaryLevel *WordDictionary, word string, curIndex int) bool {
	if curDictionaryLevel == nil {
		return false
	}
	if curIndex == len(word) {
		return curDictionaryLevel.IsWord
	}
	curChar := word[curIndex]
	if curChar == '.' {
		for _, next := range curDictionaryLevel.NextChar {
			if doSearch(next, word, curIndex+1) {
				return true
			}
		}
		return false
	}
	return doSearch(curDictionaryLevel.NextChar[word[curIndex]-'a'], word, curIndex+1)
}

func (this *WordDictionary) Search(word string) bool {
	return doSearch(this, word, 0)
}

type Trie struct {
	IsWord   bool
	Children [26]*Trie
}

func NewTrie() *Trie {
	return &Trie{IsWord: false, Children: [26]*Trie{}}
}

func (this *Trie) Insert(word string) {
	curTrie := this
	for _, w := range word {
		if curTrie.Children[w-'a'] == nil {
			curTrie.Children[w-'a'] = NewTrie()
		}
		curTrie = curTrie.Children[w-'a']
	}
	curTrie.IsWord = true
}

func findWords(board [][]byte, words []string) []string {
	tire := NewTrie()
	for _, word := range words {
		tire.Insert(word)
	}

	m, n := len(board), len(board[0])

	result := make([]string, 0)
	var dfs func(curI, curJ int, curTire *Trie, curStr []byte)
	dfs = func(curI, curJ int, curTire *Trie, curStr []byte) {
		c := board[curI][curJ]
		if c == '#' || curTire.Children[c-'a'] == nil { // c == '#' means board cell has been visited
			return
		}
		curStr = append(curStr, c)
		curTire = curTire.Children[c-'a']
		if curTire.IsWord {
			result = append(result, string(curStr))
			curTire.IsWord = false
		}

		board[curI][curJ] = '#' // mark as visited

		// search left
		if curJ-1 >= 0 {
			dfs(curI, curJ-1, curTire, curStr)
		}

		// search right
		if curJ+1 < n {
			dfs(curI, curJ+1, curTire, curStr)
		}

		// search up
		if curI-1 >= 0 {
			dfs(curI-1, curJ, curTire, curStr)
		}

		// search down
		if curI+1 < m {
			dfs(curI+1, curJ, curTire, curStr)
		}

		board[curI][curJ] = c
		curStr = curStr[0 : len(curStr)-1]
	}

	curStr := make([]byte, 0, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, tire, curStr)
		}
	}

	return result
}

func shortestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	const mod = 1000000007
	base := int64(1)
	maxPalindromeEnd := 0

	// treating a string as a 26 base number
	// hash is the 26 base number value from s[0] to s[i]
	// reverseHash is the 26 base number value from s[i] to s[0]
	hash, reverseHash := int64(0), int64(0)
	for i := 0; i < n; i++ {
		v := int64(s[i] - 'a')
		hash = (hash*26 + v) % mod //
		reverseHash = (reverseHash + base*v) % mod
		base = (base * 26) % mod

		if hash == reverseHash {
			maxPalindromeEnd = i
		}
	}

	ans := make([]byte, n-1-maxPalindromeEnd+n)
	copy(ans[n-1-maxPalindromeEnd:], s[:])
	for j := 0; j < n-1-maxPalindromeEnd; j++ {
		ans[j] = s[n-1-j]
	}
	return string(ans)
}

type MinHeap struct {
	Nodes []int
}

func (mh *MinHeap) adjustDown(startIndex int) {
	i, j := startIndex, startIndex*2+1
	n := len(mh.Nodes)
	for j < n {
		if j+1 < n && mh.Nodes[j+1] < mh.Nodes[j] { // if right child exist and its value is smaller than left value
			j = j + 1
		}
		if mh.Nodes[j] >= mh.Nodes[i] { // if the smallest child is greater than its father, no need to adjust
			break
		}
		// if the smallest child is smaller than its father, move the child up
		mh.Nodes[j], mh.Nodes[i] = mh.Nodes[i], mh.Nodes[j] // swap
		i = j
		j = i*2 + 1
	}
}

func NewMinHeap(values []int) *MinHeap {
	mh := &MinHeap{Nodes: values}
	k := len(values)
	// start from the last node with children
	for i := (k - 1) / 2; i >= 0; i-- {
		mh.adjustDown(i)
	}
	return mh
}

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	minHeap := NewMinHeap(nums[0:k])
	for i := k; i < n; i++ {
		if nums[i] > minHeap.Nodes[0] {
			minHeap.Nodes[0] = nums[i]
			minHeap.adjustDown(0)
		}
	}
	return minHeap.Nodes[0]
}

func combinationSum3(k int, n int) [][]int {
	result := make([][]int, 0)
	visitInfo := map[int]bool{}

	var doCombine func(curList []int, curSum int, curI int)

	doCombine = func(curList []int, curSum int, curI int) {
		if len(curList) == k && curSum == n {
			oneResult := make([]int, len(curList))
			copy(oneResult, curList)
			result = append(result, oneResult)
			return
		}
		if len(curList) > k {
			return
		}
		if curSum > n {
			return
		}

		for i := curI; i <= 9; i++ {
			if !visitInfo[i] {
				visitInfo[i] = true
				doCombine(append(curList, i), curSum+i, i+1)
				visitInfo[i] = false
			}
		}
	}

	doCombine([]int{}, 0, 1)
	return result
}

func containsDuplicate(nums []int) bool {
	numCount := map[int]bool{}
	for _, num := range nums {
		if numCount[num] {
			return true
		}
		numCount[num] = true
	}
	return false
}

type HeapNode struct {
	x             int
	height        int
	buildingIndex int
	isEnding      bool
	heapIndex     int
}

type HeapNodeList []*HeapNode

func (h HeapNodeList) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].heapIndex, h[j].heapIndex = i, j
}

func (h HeapNodeList) Len() int {
	return len(h)
}

func (h HeapNodeList) Less(i, j int) bool {
	return h[i].height > h[j].height
}

func (h *HeapNodeList) Push(x interface{}) {
	*h = append(*h, x.(*HeapNode))
}

func (h *HeapNodeList) Pop() interface{} {
	ret := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return ret
}

func getSkyline(buildings [][]int) [][]int {
	ans := make([][]int, 0)
	heapNodeList := &HeapNodeList{}
	heap.Init(heapNodeList)
	points := make([]*HeapNode, len(buildings)*2)
	for i, building := range buildings {
		points[i*2] = &HeapNode{
			x:             building[0],
			height:        building[2],
			buildingIndex: i,
			isEnding:      false,
			heapIndex:     0, // not in a heap yet
		}
		points[i*2+1] = &HeapNode{
			x:             building[1],
			height:        building[2],
			buildingIndex: i,
			isEnding:      true,
			heapIndex:     0, // not in a heap yet
		}
	}
	sort.Slice(points, func(i, j int) bool {
		if points[i].x != points[j].x {
			return points[i].x < points[j].x
		} else if points[i].isEnding && points[j].isEnding {
			return points[i].height < points[j].height
		} else if !points[i].isEnding && !points[j].isEnding {
			return points[i].height > points[j].height
		} else {
			return !points[i].isEnding
		}
	})

	heap.Push(heapNodeList, &HeapNode{x: 0, height: 0})
	ongoingHeight := 0
	pointBuildingIndexToPoint := make([]*HeapNode, len(buildings))
	for _, point := range points {
		if !point.isEnding {
			point.heapIndex = heapNodeList.Len()
			heap.Push(heapNodeList, point) // heapIndex will be automatically modified during the swap process of the heap
			pointBuildingIndexToPoint[point.buildingIndex] = point
		} else {
			// pointBuildingIndexToPoint[point.buildingIndex] correspond to the start point if this building
			heap.Remove(heapNodeList, pointBuildingIndexToPoint[point.buildingIndex].heapIndex)
		}

		topHeight := (*heapNodeList)[0].height
		if ongoingHeight != topHeight {
			ongoingHeight = topHeight
			ans = append(ans, []int{point.x, ongoingHeight})
		}

	}
	return ans
}

func containsNearbyDuplicate(nums []int, k int) bool {
	numIndexMap := make(map[int]int)
	for i, num := range nums {
		if index, ok := numIndexMap[num]; ok && i-index <= k {
			return true
		}
		numIndexMap[num] = i
	}
	return false
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	getBucketIndex := func(v, bucketSize int) int {
		if v < 0 { // the negative number bucket start from -1,
			// for example, number range in [-bucketSize + 1, -1] will put in the first negative bucket with index -1
			return (v+1)/bucketSize - 1
		}
		return v / bucketSize
	}

	buckets := make(map[int]int) // there will only be indexDiff number of buckets
	bucketSize := valueDiff + 1
	for i, num := range nums {
		bucketIndex := getBucketIndex(num, bucketSize)
		if _, ok := buckets[bucketIndex]; ok { // if the bucketIndex already has a number
			return true
		}
		if _, ok := buckets[bucketIndex-1]; ok && abs(num-buckets[bucketIndex-1]) <= valueDiff {
			return true
		}
		if _, ok := buckets[bucketIndex+1]; ok && abs(num-buckets[bucketIndex+1]) <= valueDiff {
			return true
		}
		buckets[bucketIndex] = num
		if i >= indexDiff {
			// make sure the all the numbers in bucket are within the indexDiff
			delete(buckets, getBucketIndex(nums[i-indexDiff], bucketSize))
		}
	}
	return false
}
