package interview_related

import "testing"

func TestLargestSetOfElementsSharedOneDigit(t *testing.T) {
	t.Log(largestSetOfElementsSharedOneDigit([]int{52, 25, 11, 52, 34, 55}))
	t.Log(largestSetOfElementsSharedOneDigit([]int{71, 23, 57, 15}))
	t.Log(largestSetOfElementsSharedOneDigit([]int{11, 33, 55}))
	t.Log(largestSetOfElementsSharedOneDigit([]int{90, 90, 90}))
}
