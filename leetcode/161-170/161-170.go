package _161_170

import (
	"math"
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

func maximumGap(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	if n == 2 {
		return abs(nums[0] - nums[1])
	}

	// retrieve min and max of array
	minV, maxV := nums[0], nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > maxV {
			maxV = nums[i]
		}
		if nums[i] < minV {
			minV = nums[i]
		}
	}

	if minV == maxV {
		return 0
	}

	averageGap := (maxV-minV)/(n-1) + 1
	bucketMin := make([]int, n-1)
	bucketMax := make([]int, n-1)

	for i := 0; i < n-1; i++ {
		bucketMin[i] = math.MaxInt
		bucketMax[i] = math.MinInt
	}

	// put nums into buckets respectively
	for _, num := range nums {
		if num == maxV || num == minV {
			continue
		}

		// bucket to put
		idx := (num - minV) / averageGap

		bucketMin[idx] = min(num, bucketMin[idx])
		bucketMax[idx] = max(num, bucketMax[idx])
	}

	maxGap := math.MinInt
	prev := minV
	for i := 0; i < n-1; i++ {
		if bucketMin[i] == math.MaxInt && bucketMax[i] == math.MinInt {
			// empty bucket
			continue
		}
		maxGap = max(maxGap, bucketMin[i]-prev)
		prev = bucketMax[i]
	}
	maxGap = max(maxGap, maxV-prev)
	return maxGap
}

func compareVersion(version1 string, version2 string) int {
	versionStrList1 := strings.Split(version1, ".")
	versionStrList2 := strings.Split(version2, ".")
	m, n := len(versionStrList1), len(versionStrList2)
	i, j := 0, 0
	for i < m && j < n {
		v1, _ := strconv.Atoi(versionStrList1[i])
		v2, _ := strconv.Atoi(versionStrList2[j])
		if v1 > v2 {
			return 1
		} else if v1 < v2 {
			return -1
		}
		i++
		j++
	}
	for i < m {
		v1, _ := strconv.Atoi(versionStrList1[i])
		if v1 != 0 {
			return 1
		}
		i++
	}
	for j < n {
		v2, _ := strconv.Atoi(versionStrList2[j])
		if v2 != 0 {
			return -1
		}
		j++
	}
	return 0
}

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	negFlag := 1
	if numerator^denominator < 0 {
		negFlag = -1
	}
	var sb strings.Builder
	remain := numerator % denominator
	if remain == 0 {
		return strconv.Itoa(numerator / denominator)
	}
	remainIndexCache := map[int]int{}
	index := 0
	for remain != 0 {
		if _, ok := remainIndexCache[remain]; ok {
			break
		}
		remainIndexCache[remain] = index
		index++
		remain *= 10
		sb.WriteString(strconv.Itoa(negFlag * remain / denominator))
		remain = remain % denominator
	}
	integerPart := numerator / denominator
	fractionPart := sb.String()
	if remain == 0 {
		if negFlag == -1 && integerPart == 0 {
			return "-" + strconv.Itoa(integerPart) + "." + fractionPart
		}
		return strconv.Itoa(integerPart) + "." + fractionPart
	}
	remainIndex := remainIndexCache[remain]
	if negFlag == -1 && integerPart == 0 {
		return "-" + strconv.Itoa(integerPart) + "." + fractionPart[:remainIndex] + "(" + fractionPart[remainIndex:] + ")"
	}
	return strconv.Itoa(integerPart) + "." + fractionPart[:remainIndex] + "(" + fractionPart[remainIndex:] + ")"
}

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		v := numbers[left] + numbers[right]
		if v == target {
			return []int{left + 1, right + 1}
		} else if v < target {
			left++
		} else {
			right++
		}
	}
	return nil
}

func convertToTitle(columnNumber int) string {
	var sb strings.Builder
	for columnNumber > 26 {
		if v := uint8(columnNumber % 26); v == 0 {
			sb.WriteByte('0')
		} else {
			sb.WriteByte('@' + v)
		}
		columnNumber /= 26
	}
	sb.WriteByte('@' + uint8(columnNumber))
	res := []byte(sb.String())
	left, right := 0, len(res)-1
	for left < right {
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}
	return string(res)
}

func majorityElement(nums []int) int {
	n := len(nums)
	numCounts := map[int]int{}
	for _, num := range nums {
		numCounts[num]++
		if numCounts[num] > n/2 {
			return num
		}
	}
	return 0
}
