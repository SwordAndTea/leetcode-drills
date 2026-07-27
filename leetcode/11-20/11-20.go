package _11_20

import (
	"math"
	"sort"
)

// Container With Most Water
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxV := 0
	area := 0
	for left < right {
		if height[left] < height[right] {
			area = height[left] * (right - left)
			left++
		} else {
			area = height[right] * (right - left)
			right--
		}

		if area > maxV {
			maxV = area
		}
	}
	return maxV
}

// leetcode problem No. 12
func intToRoman(num int) string {
	ans := []rune{}
	for num/1000 > 0 {
		ans = append(ans, 'M')
		num -= 1000
	}
	if num >= 900 {
		ans = append(ans, 'C', 'M')
		num -= 900
	}
	for num/500 > 0 {
		ans = append(ans, 'D')
		num -= 500
	}
	if num >= 400 {
		ans = append(ans, 'C', 'D')
		num -= 400
	}
	for num/100 > 0 {
		ans = append(ans, 'C')
		num -= 100
	}
	if num >= 90 {
		ans = append(ans, 'X', 'C')
		num -= 90
	}
	for num/50 > 0 {
		ans = append(ans, 'L')
		num -= 50
	}
	if num >= 40 {
		ans = append(ans, 'X', 'L')
		num -= 40
	}
	for num/10 > 0 {
		ans = append(ans, 'X')
		num -= 10
	}
	if num >= 9 {
		ans = append(ans, 'I', 'X')
		num -= 9
	}
	for num/5 > 0 {
		ans = append(ans, 'V')
		num -= 5
	}
	if num >= 4 {
		ans = append(ans, 'I', 'V')
		num -= 4
	}
	for num/1 > 0 {
		ans = append(ans, 'I')
		num -= 1
	}
	return string(ans)
}

// leetcode problem No. 13
func romanToInt(s string) int {
	ans := 0
	for i, c := range s {
		if c == 'I' {
			ans += 1
		} else if c == 'V' {
			if i > 0 && s[i-1] == 'I' {
				ans -= 2
			}
			ans += 5
		} else if c == 'X' {
			if i > 0 && s[i-1] == 'I' {
				ans -= 2
			}
			ans += 10
		} else if c == 'L' {
			if i > 0 && s[i-1] == 'X' {
				ans -= 20
			}
			ans += 50
		} else if c == 'C' {
			if i > 0 && s[i-1] == 'X' {
				ans -= 20
			}
			ans += 100
		} else if c == 'D' {
			if i > 0 && s[i-1] == 'C' {
				ans -= 200
			}
			ans += 500
		} else if c == 'M' {
			if i > 0 && s[i-1] == 'C' {
				ans -= 200
			}
			ans += 1000
		}
	}
	return ans
}

// leetcode problem No. 14

func commonPrefix(s1, s2 string) string {
	i := 0
	for ; i < len(s1) && i < len(s2) && s1[i] == s2[i]; i++ {
	}
	return s1[0:i]
}

func longestCommonPrefix(strs []string) string {
	common := strs[0]
	for i := 1; i < len(strs); i++ {
		if common == "" {
			return common
		}
		common = commonPrefix(common, strs[i])
	}
	return common
}

// 3Sum
func threeSum(nums []int) [][]int {
	result := make([][]int, 0, 128)
	indexMap := make(map[int][]int)
	for i, n := range nums {
		indexes := indexMap[n]
		if indexes == nil {
			indexes = make([]int, 0, 16)
		}
		indexMap[n] = append(indexes, i)
	}
	keys := make([]int, 0, len(indexMap))
	for k, _ := range indexMap {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for i := 0; i < len(keys); i++ {
		v1 := keys[i]
		indices1 := indexMap[v1]
		for j := i; j < len(keys); j++ {
			v2 := keys[j]
			if v1 == 0 && v2 == 0 && len(indices1) <= 2 {
				continue
			}
			indices2 := indexMap[v2]
			v3 := -(v1 + v2)
			indices3 := indexMap[v3]
			if len(indices3) > 0 && v3 >= v2 {
				if v1 == v2 {
					if len(indices1) > 1 {
						result = append(result, []int{v1, v2, v3})
					}
				} else if v1 == v3 {
					if len(indices1) > 1 {
						result = append(result, []int{v1, v2, v3})
					}
				} else if v2 == v3 {
					if len(indices2) > 1 {
						result = append(result, []int{v1, v2, v3})
					}
				} else {
					result = append(result, []int{v1, v2, v3})
				}
			}
		}
	}
	return result
}

func threeSumImpl2(nums []int) [][]int {
	result := make([][]int, 0, 128)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1

		for left < right {
			if nums[i]+nums[left]+nums[right] == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if nums[i]+nums[left]+nums[right] > 0 {
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				right--
			} else {
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++
			}
		}
	}
	return result
}

// 3Sum Closest
func abs(v int) int {
	if v >= 0 {
		return v
	}
	return -v
}

func threeSumClosest(nums []int, target int) int {
	distance := math.MaxInt
	result := 0

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				return target
			} else if sum > target {
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				right--
			} else {
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++
			}

			if abs(sum-target) < distance {
				distance = abs(sum - target)
				result = sum
			}
		}
	}
	return result
}

// Letter Combinations of a Phone Number

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	digitalMap := map[uint8][]uint8{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	resultLen := 1
	for i := 0; i < len(digits); i++ {
		resultLen *= len(digitalMap[digits[i]])
	}
	result := make([]string, resultLen)

	step := resultLen
	for i := 0; i < len(digits); i++ {
		chars := digitalMap[digits[i]]
		step /= len(chars)
		charIndex := 0
		for j := 0; j < resultLen; j += step {
			char := chars[charIndex]
			for k := j; k < j+step; k++ {
				result[k] = string(append([]byte(result[k]), char))
			}
			charIndex = (charIndex + 1) % len(chars)
		}
	}

	return result
}

// 4Sum
func fourSum(nums []int, target int) [][]int {
	result := make([][]int, 0, 128)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 && nums[i] > target {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		newTarget := target - nums[i]

		for j := i + 1; j < len(nums); j++ {
			if nums[j] > 0 && nums[j] > newTarget {
				break
			}
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left := j + 1
			right := len(nums) - 1

			for left < right {
				sum := nums[j] + nums[left] + nums[right]
				if sum == newTarget {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if sum > newTarget {
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					right--
				} else {
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					left++
				}
			}
		}

	}
	return result
}

//Remove Nth Node From End of List

type ListNode struct {
	Val  int
	Next *ListNode
}

// Definition for singly-linked list.
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodes := make([]*ListNode, 30)
	h := head
	top := -1
	for h != nil {
		top++
		nodes[top] = h
		h = h.Next
	}

	nodeToMove := top - n + 1
	if nodeToMove == 0 {
		if top == 0 {
			return nil
		}
		return nodes[1]
	}

	pre := nodes[top-n]
	pre.Next = nodes[nodeToMove].Next
	return nodes[0]
}

// valid parentheses
func isValid(s string) bool {
	parenthesesPariMap := map[uint8]uint8{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := make([]uint8, 0, 128)
	for i := 0; i < len(s); i++ {
		c := s[i]
		pair := parenthesesPariMap[c]
		if pair == 0 {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != pair {
				return false
			}
			stack = stack[0 : len(stack)-1]
		}
	}

	return len(stack) == 0
}
