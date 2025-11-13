package _21_130

import "math"

func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfitV := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if v := prices[i] - minPrice; v > maxProfitV {
			maxProfitV = v
		}
	}
	return maxProfitV
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit2(prices []int) int {
	maxProfitV := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfitV += prices[i] - prices[i-1]
		}
	}
	return maxProfitV
}

func maxProfit3(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}

	// one transaction profit that sells at day i
	firstTransactionProfit := make([]int, n)
	minPrice := prices[0]
	for i := 1; i < n; i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			firstTransactionProfit[i] = prices[i] - minPrice
		}
	}

	// one transaction profit that buys at or after day i
	secondTransactionProfit := make([]int, n)
	maxPrice := prices[n-1]
	for i := n - 2; i >= 0; i-- {
		if v := maxPrice - prices[i]; v > secondTransactionProfit[i+1] {
			secondTransactionProfit[i] = v
		} else {
			secondTransactionProfit[i] = secondTransactionProfit[i+1]
		}
		if prices[i] > maxPrice {
			maxPrice = prices[i]
		}
	}
	maxProfitV := 0
	for i := 0; i < n; i++ {
		if v := firstTransactionProfit[i] + secondTransactionProfit[i]; v > maxProfitV {
			maxProfitV = v
		}
	}
	return maxProfitV
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxV := math.MinInt

	var solve func(node *TreeNode) (int, int)
	solve = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		total := node.Val
		ll, lr := solve(node.Left)
		leftMax := max(ll, lr)
		left := node.Val
		if leftMax > 0 {
			left += leftMax
			total += leftMax
		}
		rl, rr := solve(node.Right)
		rightMax := max(rl, rr)
		right := node.Val
		if rightMax > 0 {
			right += rightMax
			total += rightMax
		}
		if total > maxV {
			maxV = total
		}
		return left, right
	}

	_, _ = solve(root)

	return maxV
}

func isPalindrome(s string) bool {
	p1, p2 := 0, len(s)-1
	var c1 uint8
	var c2 uint8
	for p1 < p2 {
		for p1 < p2 && !((s[p1] >= '0' && s[p1] <= '9') || (s[p1] >= 'A' && s[p1] <= 'Z') || (s[p1] >= 'a' && s[p1] <= 'z')) {
			p1++
		}
		for p1 < p2 && !((s[p2] >= '0' && s[p2] <= '9') || (s[p2] >= 'A' && s[p2] <= 'Z') || (s[p2] >= 'a' && s[p2] <= 'z')) {
			p2--
		}
		if p1 < p2 {
			if s[p1] >= 'A' && s[p1] <= 'Z' {
				c1 = s[p1] + 32
			} else {
				c1 = s[p1]
			}

			if s[p2] >= 'A' && s[p2] <= 'Z' {
				c2 = s[p2] + 32
			} else {
				c2 = s[p2]
			}

			if c1 != c2 {
				return false
			}
			p1++
			p2--
		}

	}
	return true
}

func isWordLadder(w1, w2 string) bool {
	diff := 0
	for i := 0; i < len(w1) && diff < 2; i++ {
		if w1[i] != w2[i] {
			diff++
		}
	}
	return diff == 1
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	distanceMap := make(map[string]int)
	distanceMap[beginWord] = 0
	queue := make([]string, 1)
	queue[0] = beginWord
	for len(queue) > 0 {
		curWord := queue[0]
		queue = queue[1:]
		for _, w := range wordList {
			if _, ok := distanceMap[w]; !ok && isWordLadder(curWord, w) {
				queue = append(queue, w)
				distanceMap[w] = distanceMap[curWord] + 1
			}
		}
	}

	result := make([][]string, 0)
	if _, ok := distanceMap[endWord]; !ok { // can not reach to endWord
		return result
	}

	var dfs func(curList []string)
	dfs = func(curList []string) {
		curWord := curList[len(curList)-1]
		if distanceMap[curWord] == 1 { // reached the word that next to beginWord
			// new result is the reverse of curList
			newResult := make([]string, len(curList)+1)
			newResult[0] = beginWord
			for i := 0; i < len(curList); i++ {
				newResult[i+1] = curList[len(curList)-1-i]
			}
			result = append(result, newResult)
			return
		}

		for _, w := range wordList {
			if _, ok := distanceMap[w]; ok && isWordLadder(curWord, w) &&
				distanceMap[curWord] == distanceMap[w]+1 { // must be reachable word and the word next to curWord
				dfs(append(curList, w))
			}
		}
	}

	dfs([]string{endWord})
	return result
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	ladderMap := make(map[string][]string)
	for _, w := range wordList {
		for i := 0; i < len(w); i++ {
			wildcard := w[:i] + "*" + w[i+1:] // replace w[i] with *
			ladderMap[wildcard] = append(ladderMap[wildcard], w)
		}
	}

	distanceMap := make(map[string]int)
	distanceMap[beginWord] = 0
	queue := make([]string, 1)
	queue[0] = beginWord
	for len(queue) > 0 {
		curWord := queue[0]
		queue = queue[1:]
		for i := 0; i < len(curWord); i++ {
			wildcard := curWord[:i] + "*" + curWord[i+1:]
			for _, w := range ladderMap[wildcard] {
				if _, ok := distanceMap[w]; !ok {
					queue = append(queue, w)
					distanceMap[w] = distanceMap[curWord] + 1
					if w == endWord {
						return distanceMap[w] + 1
					}
				}
			}
		}
	}

	return 0
}

func longestConsecutive(nums []int) int {
	numInfo := make(map[int]bool)
	for _, num := range nums {
		numInfo[num] = true
	}
	sequenceLength := make(map[int]int)
	maxLength := 0
	for num, _ := range numInfo {
		curLength := 1
		next := num + 1
		for numInfo[next] {
			if sequenceLength[next] != 0 {
				curLength += sequenceLength[next]
				break
			} else {
				next++
				curLength++
			}
		}
		sequenceLength[num] = curLength
		if curLength > maxLength {
			maxLength = curLength
		}
	}
	return maxLength
}

func sumNumbers(root *TreeNode) int {
	total := 0
	var solve func(node *TreeNode, cur int)
	solve = func(node *TreeNode, cur int) {
		cur = cur*10 + node.Val
		if node.Left == nil && node.Right == nil { // leaf node
			total += cur
			return
		}
		if node.Left != nil {
			solve(node.Left, cur)
		}
		if node.Right != nil {
			solve(node.Right, cur)
		}
	}
	solve(root, 0)
	return total
}

type coordinate struct {
	i int
	j int
}

func solve(board [][]byte) {
	m, n := len(board), len(board[0])
	visitInfo := make([]bool, m*n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && !visitInfo[i*n+j] {
				visitInfo[i*n+j] = true
				queue := make([]*coordinate, 1)
				queue[0] = &coordinate{i, j}
				front := 0
				shouldCover := true
				for front < len(queue) {
					cur := queue[front]
					front++
					if cur.i == 0 || cur.j == 0 || cur.i == m-1 || cur.j == n-1 {
						shouldCover = false
					}
					// go left
					nextJ := cur.j - 1
					if nextJ >= 0 && board[cur.i][nextJ] == 'O' && !visitInfo[cur.i*n+nextJ] {
						queue = append(queue, &coordinate{cur.i, nextJ})
						visitInfo[cur.i*n+nextJ] = true
					}

					// go right
					nextJ = cur.j + 1
					if nextJ < n && board[cur.i][nextJ] == 'O' && !visitInfo[cur.i*n+nextJ] {
						queue = append(queue, &coordinate{cur.i, nextJ})
						visitInfo[cur.i*n+nextJ] = true
					}

					// go up
					nextI := cur.i - 1
					if nextI >= 0 && board[nextI][cur.j] == 'O' && !visitInfo[nextI*n+cur.j] {
						queue = append(queue, &coordinate{nextI, cur.j})
						visitInfo[nextI*n+cur.j] = true
					}

					// go down
					nextI = cur.i + 1
					if nextI < m && board[nextI][cur.j] == 'O' && !visitInfo[nextI*n+cur.j] {
						queue = append(queue, &coordinate{nextI, cur.j})
						visitInfo[nextI*n+cur.j] = true
					}
				}
				if shouldCover {
					for _, v := range queue {
						board[v.i][v.j] = 'X'
					}
				}
			}
		}
	}
}
