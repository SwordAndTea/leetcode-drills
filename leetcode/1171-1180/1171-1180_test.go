package _1171_1180

import "testing"

func TestRemoveZeroSumSublists(t *testing.T) {
	res := removeZeroSumSublists(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: -3,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
		},
	})

	t.Logf("%+v", res)
}
