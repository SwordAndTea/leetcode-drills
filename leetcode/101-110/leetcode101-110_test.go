package _01_110

import "testing"

func Test_isSymmetric(t *testing.T) {
	t.Log(isSymmetric(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 3},
		},
	}))
}

func Test_sortedArrayToBST(t *testing.T) {
	tree := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	t.Log(tree)
}
