package _1171_1180

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode problem No. 1171

func removeZeroSumSublists(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}

	cur := head
	prefixSum := 0
	prefixSumMap := map[int]*ListNode{}
	prefixSumMap[0] = dummy

	for cur != nil {
		prefixSum += cur.Val
		if node, ok := prefixSumMap[prefixSum]; ok {
			nodeToDelete := node.Next
			sumToNodeToDelete := prefixSum
			for nodeToDelete != cur { // mark node to cur as deleted
				sumToNodeToDelete += nodeToDelete.Val
				delete(prefixSumMap, sumToNodeToDelete) // important
				nodeToDelete = nodeToDelete.Next
			}
			node.Next = cur.Next
		} else {
			prefixSumMap[prefixSum] = cur
		}
		cur = cur.Next
	}

	return dummy.Next
}
