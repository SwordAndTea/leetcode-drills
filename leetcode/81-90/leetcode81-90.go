package _1_90

import "sort"

func search(nums []int, target int) bool {
	n := len(nums)

	var searchPivot func(left, right int) int
	searchPivot = func(left, right int) int {
		for left <= right {
			mid := (left + right) / 2
			if mid < n-1 && nums[mid] > nums[mid+1] {
				return mid
			}
			if mid > 0 && nums[mid-1] > nums[mid] {
				return mid - 1
			}

			if nums[mid] > nums[n-1] || nums[mid] > nums[0] {
				left = mid + 1
			} else if nums[mid] < nums[n-1] || nums[mid] < nums[0] {
				right = mid - 1
			} else {
				// search left part
				p1 := searchPivot(left, mid-1)
				if p1 != -1 {
					return p1
				}
				p2 := searchPivot(mid+1, right)
				if p2 != -1 {
					return p2
				}
				return -1
			}
		}

		return -1
	}

	pivotIndex := searchPivot(0, n-1)

	if pivotIndex != -1 && nums[pivotIndex] == target {
		return true
	}

	searchTarget := func(left, right int) bool {
		for left <= right {
			mid := (left + right) / 2
			if nums[mid] == target {
				return true
			}
			if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		return false
	}

	if pivotIndex == -1 {
		return searchTarget(0, n-1)
	}

	if target > nums[n-1] {
		return searchTarget(0, pivotIndex)
	}

	return searchTarget(pivotIndex+1, n-1)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre, cur := head, head.Next
	tmpHead := &ListNode{
		Next: head,
	}
	writeNode := tmpHead
	for cur != nil {
		if cur.Val == pre.Val {
			for cur != nil && cur.Val == pre.Val {
				cur = cur.Next
			}
			writeNode.Next = cur
			pre = cur
			if cur != nil {
				cur = cur.Next
			}
		} else {
			pre = cur
			cur = cur.Next
			writeNode = writeNode.Next
		}
	}

	return tmpHead.Next
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre, cur := head, head.Next

	for cur != nil {
		if cur.Val == pre.Val {
			for cur != nil && cur.Val == pre.Val {
				cur = cur.Next
			}
			pre.Next = cur
		}
		pre = cur
		if cur != nil {
			cur = cur.Next
		}
	}

	return head
}

func largestRectangleArea(heights []int) int {
	stack := make([]int, 0)
	i := 0
	max := 0
	// hint: find the max bottom length that with height[i]
	for i < len(heights) {
		if len(stack) == 0 || heights[i] >= heights[stack[len(stack)-1]] {
			// if current height is bigger than the top, the the bottom must start with i
			stack = append(stack, i)
			i++
		} else {
			for len(stack) != 0 && heights[i] < heights[stack[len(stack)-1]] {
				j := stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]
				bottom := 0
				if len(stack) == 0 {
					bottom = i
				} else {
					bottom = i - 1 - stack[len(stack)-1] // from i-1 to stack.top()
				}
				newArea := heights[j] * bottom
				if newArea > max {
					max = newArea
				}
			}
			stack = append(stack, i)
			i++
		}
	}

	for len(stack) != 0 {
		j := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		bottom := 0
		if len(stack) == 0 {
			bottom = i
		} else {
			bottom = i - 1 - stack[len(stack)-1] // from i-1 to stack.top()
		}
		newArea := heights[j] * bottom
		if newArea > max {
			max = newArea
		}
	}

	return max
}

func maximalRectangle(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])

	heights := make([]int, n)
	left := make([]int, n)
	right := make([]int, n)
	for j := 0; j < n; j++ {
		right[j] = n
	}

	max := 0
	for i := 0; i < m; i++ {
		leftest := 0
		rightest := n - 1
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				heights[j]++

				if leftest > left[j] {
					left[j] = leftest
				}
			} else {
				heights[j] = 0

				left[j] = 0
				leftest = j + 1
			}

			if matrix[i][n-1-j] == '1' {
				if rightest < right[n-1-j] {
					right[n-1-j] = rightest
				}
			} else {
				right[n-1-j] = n
				rightest = n - 1 - j - 1
			}

		}

		for j := 0; j < n; j++ {
			area := (right[j] - left[j] + 1) * heights[j]
			if area > max {
				max = area
			}
		}
	}

	return max
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	tmpHead := &ListNode{Next: head}
	writeP := tmpHead
	readP := head
	pre := tmpHead

	for readP != nil && readP.Val < x {
		writeP = readP
		pre = readP
		readP = readP.Next
	}

	for readP != nil {
		if readP.Val < x {
			p2 := readP
			for p2.Next != nil && p2.Next.Val < x {
				p2 = p2.Next
			}
			tmp := writeP.Next
			writeP.Next = readP
			pre.Next = p2.Next
			p2.Next = tmp
			readP = pre.Next
			writeP = p2
		} else {
			pre = readP
			readP = readP.Next
		}
	}
	return tmpHead.Next
}

var memo = map[string]bool{}

func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}

	n := len(s1)
	key := s1 + s2
	if val, ok := memo[key]; ok {
		return val
	}

	hash1, hash2, hash3 := [26]byte{}, [26]byte{}, [26]byte{}
	for i := 0; i < n-1; i++ {
		hash1[s1[i]-'a'] += 1
		hash2[s2[i]-'a'] += 1
		if hash1 == hash2 {
			s1Left := s1[0 : i+1]
			s2Left := s2[0 : i+1]

			s1Right := s1[i+1 : n]
			s2Right := s2[i+1 : n]
			if isScramble(s1Left, s2Left) && isScramble(s1Right, s2Right) {
				memo[key] = true
				return true
			}
		}
		hash3[s2[n-1-i]-'a'] += 1
		if hash1 == hash3 {
			s1Left := s1[0 : i+1]
			s2Right := s2[n-1-i : n]

			s1Right := s1[i+1 : n]
			s2Left := s2[0 : n-1-i]
			if isScramble(s1Left, s2Right) && isScramble(s1Right, s2Left) {
				memo[key] = true
				return true
			}
		}
		memo[key] = false
	}

	return false
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := 0, 0
	writeIndex := m
	num1Len := m + n
	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			nums1[writeIndex%num1Len] = nums1[i]
			i++
			writeIndex++
		} else {
			nums1[writeIndex%num1Len] = nums2[j]
			j++
			writeIndex++
		}
	}

	for i < m {
		nums1[writeIndex%num1Len] = nums1[i]
		i++
		writeIndex++
	}

	for j < n {
		nums1[writeIndex%num1Len] = nums2[j]
		j++
		writeIndex++
	}

	tmp := make([]int, m)
	copy(tmp, nums1[0:m])
	copy(nums1[0:n], nums1[m:m+n])
	copy(nums1[n:m+n], tmp)
}

func grayCode(n int) []int {
	result := make([]int, 1<<n)

	var solveGrayCode func(n int)

	solveGrayCode = func(n int) {
		if n == 1 {
			result[0] = 0
			result[1] = 1
			return
		}

		solveGrayCode(n - 1)
		leftPartEndIndex := 1 << (n - 1)
		rightPartEndIndex := 1 << n
		for i := leftPartEndIndex; i < rightPartEndIndex; i++ {
			result[i] = result[rightPartEndIndex-1-i] + leftPartEndIndex
		}
	}

	solveGrayCode(n)

	return result
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	result := make([][]int, 0)

	var solve func(curI int, currentList []int)

	solve = func(curI int, currentList []int) {
		if len(currentList) <= len(nums) {
			newResult := make([]int, len(currentList))
			copy(newResult, currentList)
			result = append(result, newResult)
		}
		for i := curI; i < len(nums); {
			// option 1: choose nums[i]
			currentList = append(currentList, nums[i])
			solve(i+1, currentList)

			// option 2: not choose nums[i]
			currentList = currentList[0 : len(currentList)-1]
			j := i + 1
			for j < len(nums) && nums[j] == nums[i] {
				j++
			}
			i = j
		}
	}

	solve(0, []int{})

	return result
}
