package interview_related

import "testing"

func TestMinCostCut(t *testing.T) {
	t.Log(minCostCut(&TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 1},
			Right: nil,
		},
		Right: &TreeNode{Val: 2},
	}))

	t.Log(minCostCut(&TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
		},
	}))
}
