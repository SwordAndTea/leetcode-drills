package _41_50

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

func trap(height []int) int {
	n := len(height)
	i := 0
	sum := 0
	//scan from begin to end;
	for i < n {
		base := height[i]
		k := i + 1
		for k < n && height[k] < base {
			k++
		}
		if k == n {
			break
		}
		i++
		for i < k {
			sum += base - height[i]
			i++
		}
	}

	// scan end to begin
	j := n - 1
	for j > i {
		base := height[j]
		k := j - 1
		for k >= i && height[k] < base {
			k--
		}
		if k < i {
			break
		}
		j--
		for j > k {
			sum += base - height[j]
			j--
		}
	}
	return sum
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

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	step := 0
	for i := 0; i < len(nums); {
		j := nums[i]
		if i+j >= len(nums)-1 {
			return step + 1
		}
		max := 0
		maxIndex := i
		for k := i + 1; k <= i+j; k++ {
			if k+nums[k] > max {
				max = nums[k] + k
				maxIndex = k
			}
		}
		i = maxIndex
		step++
	}

	return step
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

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	powMap := make(map[int64]float64)
	nn := int64(n)
	if n < 0 {
		nn = -nn
	}
	result := float64(1)
	pow := int64(2)
	cur := x
	for nn > 1 {
		if pow > nn {
			pow = 2
			cur = x
		}
		if powMap[pow] == 0 {
			cur *= cur
			powMap[pow] = cur
		}
		result *= powMap[pow]
		nn -= pow
		pow *= 2
	}

	if nn == 1 {
		result *= x
	}

	if n < 0 {
		return 1 / result
	}

	return result
}
