package _171_180

import (
	"math"
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

func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	n := len(dungeon[0])
	// hp[i][j] represents the min hp needed at position (i, j)
	// Add dummy row and column at bottom and right side
	hp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		hp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			hp[i][j] = math.MaxInt
		}
	}
	hp[m][n-1] = 1
	hp[m-1][n] = 1
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			need := min(hp[i+1][j], hp[i][j+1]) - dungeon[i][j]
			if need <= 0 {
				hp[i][j] = 1
			} else {
				hp[i][j] = need
			}
		}
	}
	return hp[0][0]
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
