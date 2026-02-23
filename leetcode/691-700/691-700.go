package _691_700

import "container/heap"

// leetcode problem No. 692

type WordFrequent struct {
	word string
	freq int
}

type WordFrequentList []*WordFrequent

func (list WordFrequentList) Len() int {
	return len(list)
}

func (list WordFrequentList) Less(i, j int) bool {
	if list[i].freq == list[j].freq {
		return list[i].word < list[j].word
	}
	return list[i].freq > list[j].freq
}

func (list WordFrequentList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list *WordFrequentList) Push(x interface{}) {
	*list = append(*list, x.(*WordFrequent))
}

func (list *WordFrequentList) Pop() interface{} {
	old := *list
	n := len(old)
	x := old[n-1]
	old[n-1] = nil
	*list = old[0 : n-1]
	return x
}

func topKFrequent(words []string, k int) []string {
	wordFreqMap := make(map[string]*WordFrequent)
	wordFreqList := make(WordFrequentList, 0)
	for _, word := range words {
		if wf, ok := wordFreqMap[word]; ok {
			wf.freq++
		} else {
			newWf := &WordFrequent{word: word, freq: 1}
			wordFreqList.Push(newWf)
			wordFreqMap[word] = newWf
		}
	}
	heap.Init(&wordFreqList)
	result := make([]string, 0, k)
	for i := 0; i < k; i++ {
		wf := heap.Pop(&wordFreqList).(*WordFrequent)
		result = append(result, wf.word)
	}
	return result

	// Follow-up: Could you solve it in O(n log(k)) time and O(n) extra space?
	// remain a heap with max k element, dynamically push or pop into that heap
}
