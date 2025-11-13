package _31_140

import "testing"

func Test_partition(t *testing.T) {
	t.Log(partition("aab"))
}

func Test_canCompleteCircuit(t *testing.T) {
	t.Log(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
}

func Test_candy(t *testing.T) {
	t.Log(candy([]int{1, 3, 2, 2, 1}))
}

func Test_copyRandomList(t *testing.T) {
	node1 := &Node2{
		Val:    3,
		Next:   nil,
		Random: nil,
	}
	node2 := &Node2{
		Val:    3,
		Next:   nil,
		Random: node1,
	}
	node3 := &Node2{
		Val:    3,
		Next:   nil,
		Random: nil,
	}
	node1.Next = node2
	node2.Next = node3
	t.Log(copyRandomList(node1))
}

func Test_wordBreak(t *testing.T) {
	t.Log(wordBreak("leetcode", []string{"leet", "code"}))
}

func Test_wordBreak2(t *testing.T) {
	t.Log(wordBreak2("leetcode", []string{"leet", "code"}))
}
