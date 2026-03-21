package interview_related

func largestSetOfElementsSharedOneDigit(nums []int) int {
	digitMap := make(map[int]int)
	for _, num := range nums {
		d1 := num % 10
		d2 := (num / 10) % 10
		if d1 == d2 {
			digitMap[d1]++
		} else {
			digitMap[d1]++
			digitMap[d2]++
		}
	}

	ans := 0
	for _, v := range digitMap {
		if v > ans {
			ans = v
		}
	}

	return ans
}
