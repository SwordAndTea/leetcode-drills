package sort

import (
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T) {
	a := []int{5, 4, 3, 2, 1}

	SelectSort(a)

	fmt.Printf("%+v", a)
}

func TestInsertSort(t *testing.T) {
	a := []int{5, 4, 3, 2, 1}

	InsertSort(a)

	fmt.Printf("%+v", a)
}

func TestBubbleSort(t *testing.T) {
	a := []int{5, 4, 3, 2, 1}

	BubbleSort(a)

	fmt.Printf("%+v", a)
}

func TestMergeSort(t *testing.T) {
	a := []int{5, 4, 3, 2, 1}

	MergeSort(a)

	fmt.Printf("%+v", a)
}

func TestQuickSort(t *testing.T) {
	a := []int{5, 3, 4, 2, 1}

	QuickSort(a)

	fmt.Printf("%+v", a)
}
