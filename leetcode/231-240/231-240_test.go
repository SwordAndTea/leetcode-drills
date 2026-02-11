package _231_240

import "testing"

func Test_isPalindrome(t *testing.T) {
	t.Log(isPalindrome(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}))
}

func Test_deleteNode(t *testing.T) {
	linkList := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}
	deleteNode(linkList.Next)
	t.Log(linkList)
}
