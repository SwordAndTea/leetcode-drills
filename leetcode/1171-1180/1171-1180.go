package _1171_1180

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode problem No. 1171

func removeZeroSumSublists(head *ListNode) *ListNode {
	prefixSumInfo := make(map[int]*ListNode)
	dummyHead := &ListNode{}
	dummyHead.Next = head
	prefixSumInfo[0] = dummyHead

	cur := head
	prefixSum := 0
	for cur != nil {
		prefixSum += cur.Val
		if previousEqualNode := prefixSumInfo[prefixSum]; previousEqualNode != nil {
			nodeToDelete := previousEqualNode.Next
			sumToNodeToDelete := prefixSum // note how to calculate the prefix sum of node after previousEqualNode
			for nodeToDelete != cur {
				sumToNodeToDelete += nodeToDelete.Val
				delete(prefixSumInfo, sumToNodeToDelete)
				nodeToDelete = nodeToDelete.Next
			}
			previousEqualNode.Next = cur.Next
		} else {
			prefixSumInfo[prefixSum] = cur
		}
		cur = cur.Next
	}

	return dummyHead.Next
}
