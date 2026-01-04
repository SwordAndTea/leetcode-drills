package _201_210

import (
	"math"
	"strings"
)

func rangeBitwiseAnd(left int, right int) int {
	for right > left {
		right &= right - 1
	}

	return right & left
}

func isHappy(n int) bool {
	numMap := map[int]bool{
		n: true,
	}

	for n != 1 {
		next := 0
		for n != 0 {
			v := n % 10
			next += v * v
			n /= 10
		}
		if numMap[next] {
			return false
		}
		numMap[next] = true
		n = next
	}
	return true
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	tmp := &ListNode{
		Val:  0,
		Next: head,
	}
	pre := tmp
	cur := head
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
			cur = cur.Next
		} else {
			pre = cur
			cur = cur.Next
		}
	}
	return tmp.Next
}

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	notPrime := make([]bool, n)
	marker := 2
	notPrime[0] = true
	notPrime[1] = true
	m := int(math.Sqrt(float64(n)))
	for marker <= m {
		next := marker * 2
		for next < n {
			notPrime[next] = true
			next += marker
		}

		// get next marker
		i := marker + 1
		for i <= m {
			if !notPrime[i] {
				break
			}
			i++
		}
		marker = i
	}

	ans := 0
	for i := 2; i < n; i++ {
		if !notPrime[i] {
			ans++
		}
	}
	return ans
}

func isIsomorphic(s string, t string) bool {
	n := len(s)
	characterMap := map[byte]byte{}
	reverseMap := map[byte]byte{}
	for i := 0; i < n; i++ {
		if v1, ok1 := characterMap[s[i]]; ok1 {
			if v1 != t[i] {
				return false
			}
		} else if v2, ok2 := reverseMap[t[i]]; ok2 {
			if v2 != s[i] {
				return false
			}
		} else {
			characterMap[s[i]] = t[i]
			reverseMap[t[i]] = s[i]
		}
	}
	return true
}

func reverseList(head *ListNode) *ListNode {
	tmpHead := &ListNode{
		Val:  0,
		Next: nil,
	}
	p := head
	for p != nil {
		tmp := p.Next
		p.Next = tmpHead.Next
		tmpHead.Next = p
		p = tmp
	}
	return tmpHead.Next
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)
	inDegrees := make(map[int]int)
	for _, p := range prerequisites {
		inDegrees[p[0]] += 1
		graph[p[1]] = append(graph[p[1]], p[0])
	}
	hasNodeToRemove := true
	for hasNodeToRemove {
		hasNodeToRemove = false
		for fromNode, _ := range graph {
			if inDegrees[fromNode] == 0 {
				hasNodeToRemove = true
				for _, toNode := range graph[fromNode] {
					inDegrees[toNode] -= 1
				}
				delete(graph, fromNode)
			}
		}
	}
	for _, v := range inDegrees {
		if v != 0 {
			return false
		}
	}
	return true
}

type Trie struct {
	wordMap map[string]bool
}

func Constructor() Trie {
	return Trie{wordMap: make(map[string]bool)}
}

func (this *Trie) Insert(word string) {
	this.wordMap[word] = true
}

func (this *Trie) Search(word string) bool {
	return this.wordMap[word]
}

func (this *Trie) StartsWith(prefix string) bool {
	for k, _ := range this.wordMap {
		if strings.HasPrefix(k, prefix) {
			return true
		}
	}
	return false
}

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 1 {
		if nums[0] >= target {
			return 1
		}
		return 0
	}
	left, right := 0, 0
	curSum := 0
	minLength := math.MaxInt
	for right < len(nums) {
		curSum += nums[right]
		if curSum >= target {
			// move left
			for left <= right && curSum-nums[left] >= target {
				curSum -= nums[left]
				left++
			}
			if v := right - left + 1; v < minLength {
				minLength = v
			}
		}
		right++
	}
	if minLength == math.MaxInt {
		return 0
	}
	return minLength
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	result := make([]int, 0, numCourses)
	graph := make(map[int][]int)
	inDegrees := make(map[int]int)
	for _, p := range prerequisites {
		inDegrees[p[0]] += 1
		graph[p[1]] = append(graph[p[1]], p[0])
	}
	visited := make(map[int]bool)
	hasNodeToRemove := true
	for hasNodeToRemove {
		hasNodeToRemove = false
		for fromNode := 0; fromNode < numCourses; fromNode++ {
			if inDegrees[fromNode] == 0 && !visited[fromNode] {
				hasNodeToRemove = true
				result = append(result, fromNode)
				visited[fromNode] = true
				for _, toNode := range graph[fromNode] {
					inDegrees[toNode] -= 1
				}
				delete(graph, fromNode)
			}
		}
	}
	if len(result) < numCourses {
		return []int{}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	dp := make([]int, n)
	// not rob last
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n-1; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	v1 := dp[n-2]

	// not rob first
	dp[0] = 0
	dp[1] = nums[1]
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return max(dp[n-1], v1)
}
