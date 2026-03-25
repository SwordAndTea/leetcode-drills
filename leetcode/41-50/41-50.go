package _41_50

import "math"

func firstMissingPositive(nums []int) int {
	numsLen := len(nums)
	i := 0
	// change nums like num[0] = 1, nums[1] = 2, ...
	for i < numsLen {
		if nums[i] <= 0 || nums[i] > numsLen || nums[nums[i]-1] == nums[i] {
			// note: the judge condition can not use nums[i] == i + 1
			i++
			continue
		}
		nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
	}

	i = 0
	for i < numsLen && nums[i] == i+1 {
		i++
	}

	return i + 1
}

// leetcode problem No. 42

func trap(height []int) int {
	// there is a common solution for this kind of problem, see the leetcode problem No. 407
	left := 0
	right := len(height) - 1
	ans := 0
	for left < right-1 {
		if height[left] < height[right] {
			ans += max(height[left]-height[left+1], 0)
			height[left+1] = max(height[left], height[left+1])
			left++
		} else {
			ans += max(height[right]-height[right-1], 0)
			height[right-1] = max(height[right], height[right-1])
			right--
		}
	}
	return ans
}

func multiply(num1 string, num2 string) string {
	len1, len2 := len(num1), len(num2)
	totalLen := len1 + len2
	result := make([]byte, totalLen)
	carry := uint8(0)
	for i := len1 - 1; i >= 0; i-- {
		v1 := num1[i] - '0'
		writeIndex := totalLen - (len1 - i)
		for j := len2 - 1; j >= 0; j-- {
			v := v1*(num2[j]-'0') + carry + result[writeIndex]
			carry = v / 10
			result[writeIndex] = v % 10
			writeIndex--
		}

		for carry != 0 {
			v := result[writeIndex] + carry
			carry = v / 10
			result[writeIndex] = v % 10
			writeIndex--
		}
	}

	for i := 0; i < totalLen; i++ {
		if result[i] != 0 {
			for j := i; j < totalLen; j++ {
				result[j] += '0'
			}
			return string(result[i:])
		}
	}

	return "0"
}

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1) // dp[i][j] means whether s[0:i) and p[0: j) match

	for i := 0; i < len(s)+1; i++ {
		dp[i] = make([]bool, len(p)+1)
		// dp[i][0] = false
	}

	dp[0][0] = true

	for i := 1; i < len(p)+1; i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-1]
		}
	}

	for i := 1; i < len(s)+1; i++ {
		for j := 1; j < len(p)+1; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else if p[j-1] == '?' || p[j-1] == s[i-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}

	return dp[len(s)][len(p)]
}

// leetcode problem No. 45

func jump(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[n-1] = 0
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j <= i+nums[i] && j < n; j++ {
			if dp[j] != math.MaxInt {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[0]
}

func permute(nums []int) [][]int {
	total := 1
	for i := 2; i <= len(nums); i++ {
		total *= i
	}

	result := make([][]int, total)

	resIndex := 0
	var solvePermute func(current []int, selectInfo map[int]bool)

	solvePermute = func(current []int, selectInfo map[int]bool) {
		if len(current) == len(nums) {
			result[resIndex] = make([]int, len(nums))
			copy(result[resIndex], current)
			resIndex++
			return
		}
		for _, v := range nums {
			if !selectInfo[v] {
				selectInfo[v] = true
				solvePermute(append(current, v), selectInfo)
				selectInfo[v] = false
			}
		}
	}

	solvePermute([]int{}, map[int]bool{})

	return result
}

func permuteUnique(nums []int) [][]int {
	result := make([][]int, 0)

	numCountMap := make(map[int]int)

	for _, v := range nums {
		numCountMap[v] += 1
	}

	var solvePermute func(current []int, selectInfo map[int]int)

	solvePermute = func(current []int, selectInfo map[int]int) {
		if len(current) == len(nums) {
			newResult := make([]int, len(nums))
			copy(newResult, current)
			result = append(result, newResult)
			return
		}
		for k, v := range numCountMap {
			if selectInfo[k] < v {
				selectInfo[k] += 1
				solvePermute(append(current, k), selectInfo)
				selectInfo[k] -= 1
			}
		}
	}

	solvePermute([]int{}, map[int]int{})

	return result
}

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < (n+1)/2; i++ {
		for j := i; j < n-i-1; j++ {
			matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j], matrix[n-1-j][i] =
				matrix[n-1-j][i], matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j]
		}
	}
}

// leetcode problem No. 49

func groupAnagrams(strs []string) [][]string {
	to26 := func(str string) [26]uint8 {
		result := [26]uint8{}
		for _, c := range str {
			result[c-'a'] += 1
		}
		return result
	}

	infoMap := make(map[[26]uint8][]string)
	for _, s := range strs {
		sBytes := to26(s)
		infoMap[sBytes] = append(infoMap[sBytes], s)
	}

	result := make([][]string, 0, len(infoMap))
	for _, v := range infoMap {
		result = append(result, v)
	}

	return result
}

// leetcode problem No. 50

func myPow(x float64, n int) float64 {
	isNegative := false
	if n < 0 {
		isNegative = true
		n = -n
	}
	ans := 1.0
	for n > 0 {
		cur := x
		base := 1
		for base*2 <= n {
			cur = cur * cur
			base *= 2
		}
		ans *= cur
		n -= base
	}

	if isNegative {
		return 1 / ans
	}
	return ans
}
