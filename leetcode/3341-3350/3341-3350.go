package _3341_3350

func numProduct(n int) int {
	if n == 0 {
		return 0
	}
	ans := 1
	for n != 0 {
		ans *= n % 10
		n /= 10
	}
	return ans
}

// leetcode problem No. 3345
func smallestNumber(n int, t int) int {
	for numProduct(n)%t != 0 {
		n++
	}
	return n
}
