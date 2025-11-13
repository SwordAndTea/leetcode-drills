package sort

func SelectSort(input []int) {
	for i := 0; i < len(input); i++ {
		minIndex := i
		for j := i + 1; j < len(input); j++ {
			if input[j] < input[minIndex] {
				minIndex = j
			}
		}
		input[i], input[minIndex] = input[minIndex], input[i]
	}
}

func InsertSort(input []int) {
	for i := 1; i < len(input); i++ {
		value := input[i]
		j := i
		for ; j > 0 && value < input[j-1]; j-- {
			input[j] = input[j-1]
		}

		input[j] = value
	}
}

func BubbleSort(input []int) {
	for i := 0; i < len(input); i++ {
		for j := 1; j < len(input)-i; j++ {
			if input[j] < input[j-1] {
				input[j], input[j-1] = input[j-1], input[j]
			}
		}
	}
}

func mergeSortImpl(input []int, start, end int, mergePlace []int) {
	if start == end {
		return
	}
	mid := (start + end) / 2
	mergeSortImpl(input, start, mid, mergePlace)
	mergeSortImpl(input, mid+1, end, mergePlace)
	// merge
	i, j, k := start, mid+1, start
	for ; i <= mid && j <= end; k++ {
		if input[i] <= input[j] {
			mergePlace[k] = input[i]
			i++
		} else {
			mergePlace[k] = input[j]
			j++
		}
	}

	for ; i <= mid; i++ {
		mergePlace[k] = input[i]
		k++
	}

	for ; j <= end; j++ {
		mergePlace[k] = input[j]
		k++
	}

	copy(input[start:end+1], mergePlace[start:end+1])
}

func MergeSort(input []int) {
	mergePlace := make([]int, len(input))
	mergeSortImpl(input, 0, len(input)-1, mergePlace)
}

func partition(input []int, left, right int) int {
	guard := input[left] // choose first value as guard

	i, j := left, right
	for i < j {
		for ; i < j && input[j] >= guard; j-- {
		}
		for ; i < j && input[i] <= guard; i++ {
		}
		input[i], input[j] = input[j], input[i]
	}
	// now, i is equal to j
	input[i], input[left] = input[left], input[i]
	return i
}

func quickSortImpl(input []int, start, end int) {
	if start < end {
		partitionIndex := partition(input, start, end)
		quickSortImpl(input, start, partitionIndex-1)
		quickSortImpl(input, partitionIndex+1, end)
	}
}

func QuickSort(input []int) {
	quickSortImpl(input, 0, len(input)-1)
}
