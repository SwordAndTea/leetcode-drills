package _691_700

import (
	"container/heap"
	"math"
)

// leetcode problem No. 692

type WordFrequent struct {
	word string
	freq int
}

type WordFrequentList []*WordFrequent

func (list WordFrequentList) Len() int {
	return len(list)
}

func (list WordFrequentList) Less(i, j int) bool {
	if list[i].freq == list[j].freq {
		return list[i].word < list[j].word
	}
	return list[i].freq > list[j].freq
}

func (list WordFrequentList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list *WordFrequentList) Push(x interface{}) {
	*list = append(*list, x.(*WordFrequent))
}

func (list *WordFrequentList) Pop() interface{} {
	old := *list
	n := len(old)
	x := old[n-1]
	old[n-1] = nil
	*list = old[0 : n-1]
	return x
}

func topKFrequent(words []string, k int) []string {
	wordFreqMap := make(map[string]*WordFrequent)
	wordFreqList := make(WordFrequentList, 0)
	for _, word := range words {
		if wf, ok := wordFreqMap[word]; ok {
			wf.freq++
		} else {
			newWf := &WordFrequent{word: word, freq: 1}
			wordFreqList.Push(newWf)
			wordFreqMap[word] = newWf
		}
	}
	heap.Init(&wordFreqList) // the time complexity of init heap is O(n)
	result := make([]string, 0, k)
	for i := 0; i < k; i++ {
		wf := heap.Pop(&wordFreqList).(*WordFrequent)
		result = append(result, wf.word)
	}
	return result

	// Follow-up: Could you solve it in O(n log(k)) time and O(n) extra space?
	// remain a min heap (note: it's a mean heap) with max k element, dynamically push or pop into that heap
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// leetcode problem No. 694
func numDistinctIslands(grid [][]int) int {
	islandHash := func(islandCells [][2]int) int {
		minX, minY := math.MaxInt32, math.MaxInt32
		for _, islandCell := range islandCells {
			if islandCell[0] < minX {
				minX = islandCell[0]
			}
			if islandCell[1] < minY {
				minY = islandCell[1]
			}
		}
		ans := 0
		for _, islandCell := range islandCells {
			ans += islandCell[0] - minX + islandCell[1] - minY
		}
		return ans
	}

	m, n := len(grid), len(grid[0])
	islands := make(map[int]int)
	step := []int{-1, 0, 1, 0, -1}
	bfs := func(x, y int) {
		queue := [][2]int{{x, y}}
		k := 0
		for k < len(queue) {
			cur := queue[k]
			k++
			for i := 0; i < len(step)-1; i++ {
				xx := cur[0] + step[i]
				yy := cur[1] + step[i+1]

				if xx >= 0 && xx < m && yy >= 0 && yy < n && grid[xx][yy] == 1 {
					queue = append(queue, [2]int{xx, yy})
					grid[xx][yy] = 0 // mark grid[xx][yy] as visited
				}
			}
		}
		hashV := islandHash(queue)
		if _, ok := islands[hashV]; !ok {
			islands[hashV] = 1
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				bfs(i, j)
			}
		}
	}
	return len(islands)
}

// leetcode problem No. 696

func countBinarySubstrings(s string) int {
	if len(s) < 2 {
		return 0
	}
	consecutiveCount := 1
	preConsecutiveCount := 0
	ans := 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			consecutiveCount++
		} else {
			ans += min(consecutiveCount, preConsecutiveCount)
			preConsecutiveCount = consecutiveCount
			consecutiveCount = 1
		}
	}

	ans += min(consecutiveCount, preConsecutiveCount)

	return ans
}

// leetcode problem No. 697

func findShortestSubArray(nums []int) int {
	freqMap := make(map[int]int)
	maxFreq := 0

	for _, num := range nums {
		freqMap[num]++
		if freqMap[num] > maxFreq {
			maxFreq = freqMap[num]
		}
	}

	maxOfNonNegMap := func(m map[int]int) int {
		maxV := 0
		for _, v := range m {
			if v > maxV {
				maxV = v
			}
		}

		return maxV
	}

	if maxFreq == 1 {
		return 1
	}

	left := 0
	right := 0
	freqMap = make(map[int]int)
	curMaxFreq := 0
	ans := math.MaxInt
	for right < len(nums) {
		for right < len(nums) && curMaxFreq < maxFreq {
			freqMap[nums[right]]++
			if freqMap[nums[right]] > curMaxFreq {
				curMaxFreq = freqMap[nums[right]]
			}
			right++
		}

		if curMaxFreq == maxFreq {
			for left < right && curMaxFreq == maxFreq {
				freqMap[nums[left]]--
				curMaxFreq = maxOfNonNegMap(freqMap)
				left++
			}

			if v := right - left + 1; v < ans {
				ans = v
			}
		}
	}

	return ans
}

// leetcode problem No. 698
func canPartitionKSubsets(nums []int, k int) bool {
	n := len(nums)
	sum := 0
	maxV := nums[0]
	for _, num := range nums {
		sum += num
		if num > maxV {
			maxV = num
		}
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k
	if target < maxV {
		return false
	}

	visited := make([]bool, n)

	// since each set is like a combination, so we use a startIndex to cut search branch
	var backtracking func(startIndex, curSum, remainK int) bool
	backtracking = func(startIndex, curSum, remainK int) bool {
		if curSum == target {
			if remainK == 1 {
				return true
			}
			return backtracking(0, 0, remainK-1)
		}

		for i := startIndex; i < n; i++ {
			if !visited[i] && curSum+nums[i] <= target {
				visited[i] = true
				if backtracking(i+1, curSum+nums[i], remainK) {
					return true
				}
				visited[i] = false
			}
		}

		return false
	}

	return backtracking(0, 0, k)
}
