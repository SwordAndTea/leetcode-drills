package _171_180

import (
	"sort"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func titleToNumber(columnTitle string) int {
	n := len(columnTitle)
	num := 0
	base := 1
	for i := n - 1; i >= 0; i-- {
		num += int(columnTitle[i]-'@') * base
		base *= 26
	}
	return num
}

func trailingZeroes(n int) int {
	if n == 0 {
		return 0
	}
	return n/5 + trailingZeroes(n/5)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode problem No. 173
type BSTIterator struct {
	NodeStack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	nodeStack := make([]*TreeNode, 0, 4)
	p := root
	for p != nil {
		nodeStack = append(nodeStack, p)
		p = p.Left
	}
	return BSTIterator{NodeStack: nodeStack}
}

func (this *BSTIterator) Next() int {
	p := this.NodeStack[len(this.NodeStack)-1]
	this.NodeStack = this.NodeStack[:len(this.NodeStack)-1]
	if p.Right != nil {
		q := p.Right
		for q != nil {
			this.NodeStack = append(this.NodeStack, q)
			q = q.Left
		}
	}
	return p.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.NodeStack) != 0
}

// leetcode problem No. 174
func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	dp := make([][]int, m) // dp[i][j] stands for the minimum hp required at cell[i][j]
	for i := range dp {
		dp[i] = make([]int, n)
	}

	if need := 1 - dungeon[m-1][n-1]; need > 0 {
		dp[m-1][n-1] = need
	} else {
		dp[m-1][n-1] = 1
	}

	for i := m - 2; i >= 0; i-- {
		if need := dp[i+1][n-1] - dungeon[i][n-1]; need > 0 {
			dp[i][n-1] = need
		} else {
			dp[i][n-1] = 1
		}
	}

	for j := n - 2; j >= 0; j-- {
		if need := dp[m-1][j+1] - dungeon[m-1][j]; need > 0 {
			dp[m-1][j] = need
		} else {
			dp[m-1][j] = 1
		}
	}

	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			need := min(dp[i+1][j], dp[i][j+1]) - dungeon[i][j]
			if need > 0 {
				dp[i][j] = need
			} else {
				dp[i][j] = 1
			}
		}
	}

	return dp[0][0]
}

func largestNumber(nums []int) string {
	var compareStr func(str1, str2 string) bool
	compareStr = func(str1, str2 string) bool {
		m, n := len(str1), len(str2)
		k, l := 0, 0
		for k < m && l < n {
			if str1[k] < str2[l] {
				return false
			} else if str1[k] > str2[l] {
				return true
			}
			k++
			l++
		}
		if k < m {
			return compareStr(str1[k:], str2)
		} else if l < n {
			return compareStr(str1, str2[l:])
		}
		return false
	}

	numStr := make([]string, len(nums))
	for i, num := range nums {
		numStr[i] = strconv.Itoa(num)
	}
	sort.Slice(numStr, func(i, j int) bool {
		return compareStr(numStr[i], numStr[j])
	})
	sb := strings.Builder{}
	for _, s := range numStr {
		sb.WriteString(s)
	}
	res := sb.String()
	i := 0
	for i < len(res)-1 {
		if res[i] != '0' {
			break
		}
		i++
	}
	return res[i:]
}
