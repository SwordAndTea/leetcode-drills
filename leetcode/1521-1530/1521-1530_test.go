package _1521_1530

import (
	"testing"
)

func TestCountPairs(t *testing.T) {
	tree := &TreeNode{
		Val: 78,
		Left: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 73,
				Left: &TreeNode{
					Val: 30,
				},
			},
			Right: &TreeNode{
				Val: 98,
				Left: &TreeNode{
					Val: 63,
				},
				Right: &TreeNode{
					Val: 32,
				},
			},
		},
		Right: &TreeNode{
			Val: 81,
			Left: &TreeNode{
				Val: 36,
			},
		},
	}
	t.Log(countPairs(tree, 6))
}
