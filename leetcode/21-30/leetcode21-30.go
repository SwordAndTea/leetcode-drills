package _1_30

import (
	"algorithm/leetcode/1-10"
	"math"
)

// merge two sorted lists
func mergeTwoLists(list1 *__10.ListNode, list2 *__10.ListNode) *__10.ListNode {
	p1, p2 := list1, list2
	head := &__10.ListNode{
		Val:  0,
		Next: nil,
	}
	cur := head
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			cur.Next = p1
			cur = p1
			p1 = p1.Next
		} else {
			cur.Next = p2
			cur = p2
			p2 = p2.Next
		}
	}

	if p1 != nil {
		cur.Next = p1
	}

	if p2 != nil {
		cur.Next = p2
	}

	return head.Next
}

var generateParenthesisResultMap = map[int][]string{
	1: {"()"},
	2: {"()()", "(())"},
	3: {"((()))", "(()())", "(())()", "()(())", "()()()"},
}

// Generate Parentheses
func generateParenthesis(n int) []string {
	if generateParenthesisResultMap[n] != nil {
		return generateParenthesisResultMap[n]
	}

	nextLevelResult := generateParenthesis(n - 1)
	resultStored := make(map[string]bool)

	newResult := make([]string, 0, len(nextLevelResult)*3)
	for _, v := range nextLevelResult {
		// insert in first place
		newStr := string(append([]byte(v), []byte("()")...))
		if !resultStored[newStr] {
			newResult = append(newResult, newStr)
			resultStored[newStr] = true
		}

		// insert in middle place
		newStr = string(append(append([]byte("("), []byte(v)...), ')'))
		if !resultStored[newStr] {
			newResult = append(newResult, newStr)
			resultStored[newStr] = true
		}

		// insert in back place
		newStr = string(append([]byte("()"), []byte(v)...))
		if !resultStored[newStr] {
			newResult = append(newResult, newStr)
			resultStored[newStr] = true
		}
	}

	for i := 2; i <= n-2; i++ {
		r1 := generateParenthesis(i)
		r2 := generateParenthesis(n - i)
		for _, v1 := range r1 {
			for _, v2 := range r2 {
				newStr := string(append([]byte(v1), []byte(v2)...))
				if !resultStored[newStr] {
					newResult = append(newResult, newStr)
					resultStored[newStr] = true
				}
			}
		}
	}

	generateParenthesisResultMap[n] = newResult
	return newResult
}

// Merge k Sorted Lists
func mergeKListImpl(lists []*__10.ListNode, start, end int) *__10.ListNode {
	if start == end {
		return lists[start]
	}
	mid := (start + end) / 2
	left := mergeKListImpl(lists, start, mid)
	right := mergeKListImpl(lists, mid+1, end)

	return mergeTwoLists(left, right)
}

func mergeKLists(lists []*__10.ListNode) *__10.ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	}
	return mergeKListImpl(lists, 0, length-1)
}

// Swap Nodes in Pairs
func swapPairs(head *__10.ListNode) *__10.ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	p1 := head
	head = head.Next
	pre := &__10.ListNode{
		Val:  0,
		Next: nil,
	}

	for p1 != nil && p1.Next != nil {
		pre.Next = p1.Next
		p1.Next = p1.Next.Next
		pre.Next.Next = p1
		pre = p1
		p1 = p1.Next
	}

	return head
}

// Reverse Nodes in k-Group
func reverseKGroup(head *__10.ListNode, k int) *__10.ListNode {
	if k == 1 {
		return head
	}

	stack := make([]*__10.ListNode, k)
	top := -1

	pre := &__10.ListNode{
		Val:  0,
		Next: nil,
	}

	newHead := pre
	for head != nil {
		for head != nil && top+1 < k {
			top++
			stack[top] = head
			head = head.Next
		}

		if top+1 < k {
			break
		}

		// do reverse for this k group
		p := pre
		for top >= 0 {
			p.Next = stack[top]
			top--
			p = p.Next
		}
		pre = p
		p.Next = head
	}

	return newHead.Next
}

// Remove Duplicates from Sorted Array
func removeDuplicates(nums []int) int {
	writeIndex, scanIndex := 1, 1

	for scanIndex < len(nums) {
		if nums[scanIndex] == nums[scanIndex-1] {
			scanIndex++
		} else {
			if writeIndex != scanIndex {
				nums[writeIndex] = nums[scanIndex]
			}
			writeIndex++
			scanIndex++
		}
	}
	return writeIndex
}

// Remove Element
func removeElement(nums []int, val int) int {
	writeIndex, scanIndex := 0, 0

	for scanIndex < len(nums) {
		if nums[scanIndex] == val {
			scanIndex++
		} else {
			if writeIndex != scanIndex {
				nums[writeIndex] = nums[scanIndex]
			}
			writeIndex++
			scanIndex++
		}
	}
	return writeIndex
}

func getNext(s string) []int {
	if len(s) == 0 {
		return nil
	}
	next := make([]int, len(s))
	next[0] = 0
	for i := 1; i < len(s); i++ {
		j := next[i-1]
		for j != 0 && s[i] != s[j] {
			j = next[j-1]
		}

		if s[i] == s[j] {
			next[i] = j + 1
		} else {
			next[i] = j // j is 0
		}
	}
	return next
}

// Find the Index of the First Occurrence in a String
func strStr(haystack string, needle string) int {
	patternNext := getNext(needle)
	i, j := 0, 0

	for i < len(haystack) {
		if haystack[i] == needle[j] {
			i++
			j++
			if j == len(needle) {
				return j - len(needle)
			}
		} else if j == 0 {
			i++
		} else {
			j = patternNext[j-1]
		}
	}

	return -1
}

func divideInt64(dividend int64, divisor int64) int64 {
	if dividend < 0 && divisor < 0 {
		return divideInt64(-dividend, -divisor)
	}
	isAllPositive := int64(1)
	if dividend < 0 {
		isAllPositive = -1
		dividend = -dividend
	}
	if divisor < 0 {
		isAllPositive = -1
		divisor = -divisor
	}
	sum := divisor
	multiplicationFactor := int64(1)
	result := int64(0)
	for dividend >= divisor {
		sum = divisor
		multiplicationFactor = 1
		for sum+sum <= dividend {
			sum = sum + sum
			multiplicationFactor += multiplicationFactor
		}
		result += multiplicationFactor
		dividend = dividend - sum
	}

	return result * isAllPositive
}

func divide(dividend int, divisor int) int {
	result := divideInt64(int64(dividend), int64(divisor))
	if result < math.MinInt32 {
		return math.MinInt32
	}
	if result > math.MaxInt32 {
		return math.MaxInt32
	}
	return int(result)
}

func findSubstring(s string, words []string) []int {
	wordsMap := make(map[string]int)

	for _, w := range words {
		wordsMap[w] += 1
	}

	wordLen := len(words[0])
	concatenatedWordsLen := len(words) * wordLen

	result := make([]int, 0, 16)

	for i := 0; i <= len(s)-concatenatedWordsLen; i++ {
		w := s[i : i+wordLen]
		if wordsMap[w] != 0 {
			wordAppearCount := make(map[string]int)
			wordAppearCount[w] = 1
			appearCount := 1
			if appearCount == len(words) {
				result = append(result, i)
				continue
			}
			strEndIndex := i + concatenatedWordsLen
			for j := i + wordLen; j < strEndIndex && j+wordLen <= strEndIndex; j += wordLen {
				w = s[j : j+wordLen]
				wordAppearCount[w] += 1
				if wordAppearCount[w] <= wordsMap[w] {
					appearCount += 1
					if appearCount == len(words) {
						result = append(result, i)
						break
					}
				} else {
					break
				}
			}
		}
	}

	return result
}
