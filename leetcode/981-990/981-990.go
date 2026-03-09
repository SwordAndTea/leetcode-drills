package _981_990

import "slices"

// leetcode problem No. 989

func addToArrayForm(num []int, k int) []int {
	kNum := []int{}
	for k != 0 {
		kNum = append(kNum, k%10)
		k /= 10
	}
	slices.Reverse(kNum)
	i := len(num) - 1
	j := len(kNum) - 1
	remain := 0

	ans := make([]int, max(i, j)+2)
	l := len(ans) - 1
	for i >= 0 && j >= 0 {
		v := num[i] + kNum[j] + remain
		ans[l] = v % 10
		remain = v / 10
		i--
		j--
		l--
	}

	for i >= 0 {
		v := num[i] + remain
		ans[l] = v % 10
		remain = v / 10
		i--
		l--
	}

	for j >= 0 {
		v := num[j] + remain
		ans[l] = v % 10
		remain = v / 10
		j--
		l--
	}
	if remain != 0 {
		ans[0] = remain
		return ans
	}
	return ans[1:]
}
