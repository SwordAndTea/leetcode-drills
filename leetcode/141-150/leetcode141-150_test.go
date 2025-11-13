package _41_150

import "testing"

func Test_reorderList(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	reorderList(head)
}

func Test_LRU(t *testing.T) {
	lru := Constructor(3)
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 3)
	lru.Put(4, 4)
	t.Log(lru.Get(4))
	t.Log(lru.Get(3))
	t.Log(lru.Get(2))
	t.Log(lru.Get(1))
	lru.Put(5, 5)
	t.Log(lru.Get(1))
	t.Log(lru.Get(2))
	t.Log(lru.Get(3))
	t.Log(lru.Get(4))
	t.Log(lru.Get(5))
}
