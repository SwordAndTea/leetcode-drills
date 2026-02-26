package _1_10

import "math"

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbersImpl(l1 *ListNode, l2 *ListNode, more int) *ListNode {
	if l1 != nil && l2 != nil {
		l1.Val = l1.Val + l2.Val + more
		more = l1.Val / 10
		l1.Val = l1.Val % 10
		next := addTwoNumbersImpl(l1.Next, l2.Next, more)
		l1.Next = next
		return l1
	} else if l1 != nil {
		l1.Val = l1.Val + more
		more = l1.Val / 10
		l1.Val = l1.Val % 10
		if more > 0 {
			next := addTwoNumbersImpl(l1.Next, l2, more)
			l1.Next = next
		}
		return l1
	} else if l2 != nil {
		l2.Val = l2.Val + more
		more = l2.Val / 10
		l2.Val = l2.Val % 10
		if more > 0 {
			next := addTwoNumbersImpl(l1, l2.Next, more)
			l2.Next = next
		}
		return l2
	} else {
		if more != 0 {
			newNode := &ListNode{
				Val:  more,
				Next: nil,
			}
			return newNode
		}
		return nil
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersImpl(l1, l2, 0)
}

// 3. Longest Substring Without Repeating Characters
func lengthOfLongestSubstring(s string) int {
	//if len(s) == 0 || len(s) == 1 {
	//	return len(s)
	//}
	//i, j := 0, 0
	//m := make(map[uint8]int)
	//ans := 0
	//c := uint8(0)
	//isMovingJ := true
	//for j < len(s) {
	//	if isMovingJ {
	//		c = s[j]
	//		m[c] += 1
	//		if m[c] > 1 {
	//			isMovingJ = false
	//			if j-i > ans {
	//				ans = j - i
	//			}
	//		} else {
	//			j++
	//		}
	//	} else {
	//		c = s[i]
	//		m[c] -= 1
	//		if m[c] == 1 {
	//			isMovingJ = true
	//			j++
	//		}
	//		i++
	//	}
	//}
	//if isMovingJ && j-i > ans {
	//	ans = j - i
	//}
	//return ans

	// better solution
	charIndex := make(map[rune]int)
	startIndex := 0 // the start index of substring without duplicate characters.
	ans := 0
	for i, c := range s {
		if idx, ok := charIndex[c]; ok && idx >= startIndex {
			// ok means c is repeated, at this moment, idx points to the last repeated character
			ans = max(ans, i-startIndex)
			startIndex = idx + 1
		} else {
			ans = max(ans, i-startIndex+1)
		}
		charIndex[c] = i
	}
	return ans
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)

	if n1 > n2 {
		return findMedianSortedArrays(nums2, nums1)
	}

	n := n1 + n2
	mid := (n + 1) / 2
	low, high := 0, n1

	for low <= high {
		mid1 := (low + high) / 2 // mid index for nums1
		mid2 := mid - mid1       // mid index for nums2

		l1, l2, r1, r2 := math.MinInt, math.MinInt, math.MaxInt, math.MaxInt
		if mid1 < n1 {
			r1 = nums1[mid1]
		}
		if mid2 < n2 {
			r2 = nums2[mid2]
		}
		if mid1 >= 1 { // mid1 - 1 >= 0
			l1 = nums1[mid1-1]
		}
		if mid2 >= 1 { // mid2 - 1 >= 0
			l2 = nums2[mid2-1]
		}

		if l1 <= r2 && l2 <= r1 {
			if n%2 == 1 { // odd
				if l1 > l2 {
					return float64(l1)
				} else {
					return float64(l2)
				} // max(l1, l2)
			} else { // even
				left := l1
				if l2 > l1 {
					left = l2
				} // max(l1, l2)
				right := r1
				if r2 < r1 {
					right = r2
				} // min(r1, r2)
				return float64(left+right) / 2
			}
		} else if l1 > r2 { // Move towards the left side of nums1
			high = mid1 - 1
		} else { // Move towards the right side of nums1
			low = mid1 + 1
		}
	}

	return 0
}

func longestPalindrome(s string) string {
	strLen := len(s)
	if strLen == 1 {
		return s
	}

	dp := make([]bool, strLen)
	maxStart, maxEnd := strLen-1, strLen-1
	j := strLen - 1
	for i := strLen - 2; i >= 0; i-- {
		j = strLen - 1
		for ; j >= i+2; j-- {
			if s[i] == s[j] && dp[j-1] {
				dp[j] = true
				if j-i > maxEnd-maxStart {
					maxStart, maxEnd = i, j
				}
			} else {
				dp[j] = false
			}
		}
		dp[j] = false // now j = i+1
		if s[i] == s[j] {
			dp[j] = true
			if 1 > maxEnd-maxStart {
				maxStart, maxEnd = i, j
			}
		}
		dp[i] = true
	}

	return s[maxStart : maxEnd+1]
}

// Zigzag Conversion
func covert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	totalStep := numRows*2 - 2
	result := make([]byte, len([]byte(s)))

	cur := 0
	for i := 0; i < len(s); i += totalStep {
		result[cur] = s[i]
		cur++
	}

	for i := 1; i < numRows-1; i++ {
		curStep := (numRows - i - 1) * 2
		for j := i; j < len(s); {
			result[cur] = s[j]
			j += curStep
			curStep = totalStep - curStep
			cur++
		}
	}

	for i := numRows - 1; i < len(s); i += totalStep {
		result[cur] = s[i]
		cur++
	}

	return string(result)
}

// Reverse Integer
func reverse(x int) int {
	stack := make([]int8, 10)
	top := 0

	for x != 0 {
		stack[top] = int8(x % 10)
		x = x / 10
		top++
	}

	var result int64

	for i := 0; i < top; i++ {
		result = result*10 + int64(stack[i])
	}

	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}

	return int(result)
}

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	result := int64(0)
	i := 0
	// skipping leading space
	if s[0] == ' ' {
		for i < len(s) && s[i] == ' ' {
			i++
		}
	}

	if i == len(s) {
		return 0
	}

	// judge op
	isPositive := int64(1)
	if s[i] == '-' {
		isPositive = -1
		i++
	} else if s[i] == '+' {
		i++
	}

	for i < len(s) {
		if s[i] >= '0' && s[i] <= '9' {
			result = result*10 + int64(s[i]-'0')*isPositive
			if result > math.MaxInt32 {
				return math.MaxInt32
			} else if result < math.MinInt32 {
				return math.MinInt32
			}
			i++
		} else {
			break
		}
	}

	if result > math.MaxInt32 {
		return math.MaxInt32
	} else if result < math.MinInt32 {
		return math.MinInt32
	}

	return int(result)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x <= 9 {
		return true
	}

	stack := make([]int8, 10)
	top := 0

	for x != 0 {
		stack[top] = int8(x % 10)
		top++
		x = x / 10
	}

	left, right := 0, top-1
	for left <= right {
		if stack[left] != stack[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// Regular Expression Matching
func isMatch(s string, p string) bool {
	if p == ".*" {
		return true
	}

	dp := make([][]bool, len(s)+1)

	for i := 0; i < len(s)+1; i++ {
		dp[i] = make([]bool, len(p)+1)
		// dp[i][0] = false
	}

	dp[0][0] = true

	for j := 1; j < len(p)+1; j++ {
		if p[j-1] == '*' { // p[0] will not be *, so there j must be grater than 2
			dp[0][j] = dp[0][j-2]
		} else {
			dp[0][j] = false
		}
	}

	for i := 1; i < len(s)+1; i++ {
		for j := 1; j < len(p)+1; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2]
				if p[j-2] == '.' || s[i-1] == p[j-2] {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else if p[j-1] == '.' || s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = false
			}
		}
	}

	return dp[len(s)][len(p)]
}
