package _131_140

import (
	"strings"
	"unsafe"
)

func partition(s string) [][]string {
	isPalindrome := func(start, end int) bool {
		for start < end {
			if s[start] != s[end] {
				return false
			}
			start++
			end--
		}
		return true
	}

	result := make([][]string, 0, 1)

	var solve func(curI int, curList []string)
	solve = func(curI int, curList []string) {
		if curI >= len(s) {
			newResult := make([]string, len(curList))
			copy(newResult, curList)
			result = append(result, newResult)
			return
		}
		for i := curI; i < len(s); i++ {
			if isPalindrome(curI, i) {
				solve(i+1, append(curList, s[curI:i+1]))
			}
		}
	}

	solve(0, []string{})
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCut(s string) int {
	n := len(s)
	if n == 1 {
		return 0
	}
	dp := make([]int, n) // dp[i] means the minimum number of cut of s[0...i]
	isPalindrome := make([][]bool, n)
	for i := 0; i < n; i++ {
		isPalindrome[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		minV := i
		for j := 0; j <= i; j++ { // note: j <= i
			if s[j] == s[i] && (j >= i-1 || isPalindrome[j+1][i-1]) { // if s[j...i] is Palindrome
				isPalindrome[j][i] = true
				if j == 0 {
					minV = 0
				} else {
					minV = min(minV, dp[j-1]+1)
				}
			}
		}
		dp[i] = minV
	}
	return dp[n-1]
}

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	cloneStore := make(map[int]*Node)
	queue := []*Node{node}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, n := range cur.Neighbors {
			if _, ok := cloneStore[n.Val]; !ok {
				cloneStore[n.Val] = &Node{
					Val:       n.Val,
					Neighbors: make([]*Node, len(n.Neighbors)),
				}
				queue = append(queue, n)
			}
		}
		if _, ok := cloneStore[cur.Val]; !ok {
			cloneStore[cur.Val] = &Node{
				Val:       cur.Val,
				Neighbors: make([]*Node, len(cur.Neighbors)),
			}
		}
		curClone := cloneStore[cur.Val]
		for i, n := range cur.Neighbors {
			curClone.Neighbors[i] = cloneStore[n.Val]
		}
	}
	return cloneStore[node.Val]
}

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	for p1 := 0; p1 < n; {
		if remain := gas[p1] - cost[p1]; remain >= 0 {
			p2 := (p1 + 1) % n
			for remain >= 0 {
				remain += gas[p2] - cost[p2]
				p2 = (p2 + 1) % n
				if p2 == p1 && remain >= 0 {
					return p1
				}
			}
			for p1 < n && (remain < 0 || gas[p1]-cost[p1] < 0) {
				remain -= gas[p1] - cost[p1]
				p1++
				if p1 == p2 {
					break
				}
			}
		} else {
			p1++
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func candy(ratings []int) int {
	n := len(ratings)
	sum := n
	i := 1
	for i < n {
		if ratings[i] == ratings[i-1] {
			i++
			continue
		}

		peak := 0
		for ratings[i] > ratings[i-1] {
			peak++
			sum += peak
			i++
			if i == n {
				return sum
			}
		}

		valley := 0
		for i < n && ratings[i] < ratings[i-1] {
			valley++
			sum += valley
			i++
		}
		sum -= min(peak, valley)
	}
	return sum
}

func singleNumber(nums []int) int {
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result ^= nums[i]
	}
	return result
}

func singleNumber2(nums []int) int {
	x1, x2, mask := 0, 0, 0
	for _, n := range nums {
		x2 ^= x1 & n
		x1 ^= n
		mask = ^(x1 & x2)
		x2 &= mask
		x1 &= mask
	}

	return x1
}

type Node2 struct {
	Val    int
	Next   *Node2
	Random *Node2
}

func copyRandomList(head *Node2) *Node2 {
	if head == nil {
		return nil
	}

	p1 := head
	tmpHead := &Node2{
		Val:    0,
		Next:   nil,
		Random: nil,
	}
	p2 := tmpHead
	newNodeInfo := map[int32]*Node2{}
	for p1 != nil {
		newNode := &Node2{
			Val:    p1.Val,
			Next:   nil,
			Random: p1.Random,
		}
		originAddr := *((*int32)(unsafe.Pointer(&p1)))
		newNodeInfo[originAddr] = newNode
		p2.Next = newNode
		p1 = p1.Next
		p2 = p2.Next
	}

	p2 = tmpHead.Next
	for p2 != nil {
		if p2.Random != nil {
			randomAddr := *((*int32)(unsafe.Pointer(&p2.Random)))
			p2.Random = newNodeInfo[randomAddr]
		}
		p2 = p2.Next
	}
	return tmpHead.Next
}

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	m := len(wordDict)
	dp := make([]bool, n)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if v := i + 1 - len(wordDict[j]); v >= 0 && s[v:i+1] == wordDict[j] {
				if v == 0 || dp[v-1] {
					dp[i] = true
					break
				}
			}
		}
	}
	return dp[n-1]
}

func wordBreak2(s string, wordDict []string) []string {
	store := make(map[int][]string)
	var solve func(curI int) []string
	solve = func(curI int) []string {
		if v, ok := store[curI]; ok {
			return v
		}
		curResult := make([]string, 0)
		for _, word := range wordDict {
			if curI+len(word) <= len(s) && s[curI:curI+len(word)] == word {
				if curI+len(word) == len(s) {
					curResult = append(curResult, word)
				} else {
					nextResult := solve(curI + len(word))
					for _, v := range nextResult {
						curResult = append(curResult, strings.Join([]string{word, v}, " "))
					}
				}

			}
		}
		store[curI] = curResult
		return curResult
	}
	return solve(0)
}
