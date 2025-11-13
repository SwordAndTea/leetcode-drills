package tree

import "testing"

func TestNewNode(t *testing.T) {
	n := NewNode(1)

	t.Logf("n is %+v", n)
}

func TestAVLNode_GetHeight(t *testing.T) {
	var n *AVLNode[int] = nil

	if n.GetHeight() != 0 {
		t.Fatal("result should be 0")
	}
}

func TestNewHuffmanTree(t *testing.T) {
	ht := NewHuffmanTree([]int{1, 2, 2, 3, 6})

	t.Log(ht.Root.Data)
}
