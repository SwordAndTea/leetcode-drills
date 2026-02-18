package tree

import "errors"

type Number interface {
	~int | ~uint | ~float64 // TODO: include all number type
}

type Node[T Number] struct {
	Data  T
	Left  *Node[T]
	Right *Node[T]
}

func NewNode[T Number](data T) *Node[T] {
	return &Node[T]{
		Data:  data,
		Left:  nil,
		Right: nil,
	}
}

type BinaryTree[T Number] struct {
	Root *Node[T]
}

func preorderImpl[T Number](node *Node[T], result *[]T) {
	if node != nil {
		*result = append(*result, node.Data)
		preorderImpl(node.Left, result)
		preorderImpl(node.Right, result)
	}
}

func (n *BinaryTree[T]) Preorder() []T {
	res := make([]T, 0, 8)
	preorderImpl(n.Root, &res)
	return res
}

func inorderImpl[T Number](node *Node[T], result *[]T) {
	if node != nil {
		inorderImpl(node.Left, result)
		*result = append(*result, node.Data)
		inorderImpl(node.Right, result)
	}
}

func (n *BinaryTree[T]) Inorder() []T {
	res := make([]T, 0, 8)
	inorderImpl(n.Root, &res)
	return res
}

func postorderImpl[T Number](node *Node[T], result *[]T) {
	if node != nil {
		postorderImpl(node.Left, result)
		postorderImpl(node.Right, result)
		*result = append(*result, node.Data)
	}
}

func (n *BinaryTree[T]) Postorder() []T {
	res := make([]T, 0, 8)
	postorderImpl(n.Root, &res)
	return res
}

func (n *BinaryTree[T]) LayerOrder() []T {
	res := make([]T, 0, 8)
	if n != nil {
		queue := make([]*Node[T], 0, 8)
		queue = append(queue, n.Root)
		i := 0
		for i != len(queue) {
			curNode := queue[i]
			res = append(res, curNode.Data)
			i += 1
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}
			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
		}
	}
	return res
}

// Reconstruct rebuild a binary tree according to preorder traverse and inorder traverse
func Reconstruct[T Number](preorderResult []T, inorderResult []T) (*BinaryTree[T], error) {
	if len(preorderResult) != len(inorderResult) {
		return nil, errors.New("two list length not the same")
	}
	if len(preorderResult) == 0 {
		return nil, nil
	}
	root := NewNode(preorderResult[0])
	rootIndex := -1
	for i, v := range inorderResult {
		if v == root.Data {
			rootIndex = i
			break
		}
	}
	if rootIndex < 0 {
		return nil, errors.New("")
	}
	left, err := Reconstruct[T](preorderResult[1:rootIndex+1], inorderResult[0:rootIndex])
	if err != nil {
		return nil, err
	}
	if left != nil {
		root.Left = left.Root
	}
	right, err := Reconstruct[T](preorderResult[1+rootIndex:], inorderResult[rootIndex+1:])
	if err != nil {
		return nil, err
	}
	if right != nil {
		root.Right = right.Root
	}
	return &BinaryTree[T]{Root: root}, nil
}

// binary search tree

type BinarySearchTree[T Number] struct {
	BinaryTree[T]
}

func bstDeleteImpl[T Number](node **Node[T], data T) {
	if node == nil || *node == nil { // not find the node to remove
		return
	}

	if (*node).Data == data { // find the node to remove
		if (*node).Left == nil && (*node).Right == nil { // leaf node, we can simply delete it
			*node = nil
		} else if (*node).Left != nil { // find the predecessor
			parent := *node
			curNode := (*node).Left
			for curNode.Right != nil {
				parent = curNode
				curNode = curNode.Right
			}
			// use predecessor to replace the node
			(*node).Data = curNode.Data
			parent.Right = curNode.Left
		} else { // (*node).Right != nil, find the successor
			parent := *node
			curNode := (*node).Right
			for curNode.Left != nil {
				parent = curNode
				curNode = curNode.Left
			}
			// use successor to replace the node
			(*node).Data = data
			parent.Left = curNode.Right
		}
	} else if (*node).Data > data {
		bstDeleteImpl(&(*node).Left, data)
	} else {
		bstDeleteImpl(&(*node).Right, data)
	}
}

func (bst *BinarySearchTree[T]) Delete(data T) {
	bstDeleteImpl(&bst.Root, data)
}

// avl

type AVLNode[T Number] struct {
	Data   T
	Height int
	Left   *AVLNode[T]
	Right  *AVLNode[T]
}

func NewAVLNode[T Number](data T) *AVLNode[T] {
	return &AVLNode[T]{
		Data:   data,
		Height: 1,
	}
}

func (n *AVLNode[T]) GetHeight() int {
	if n == nil {
		return 0
	}
	return n.Height
}

func (n *AVLNode[T]) GetBalanceFactor() int {
	if n == nil {
		return 0
	}
	return n.Left.GetHeight() - n.Right.GetHeight()
}

func (n *AVLNode[T]) UpdateHeight() {
	lHeight := n.Left.GetHeight()
	rHeight := n.Right.GetHeight()
	if lHeight >= rHeight {
		n.Height = lHeight
	} else {
		n.Height = rHeight
	}
}

type AVL[T Number] struct {
	Root *AVLNode[T]
}

func leftRotation[T Number](node *AVLNode[T]) {
	if node != nil && node.Right != nil {
		right := node.Right
		node.Right = right.Left
		right.Left = node
		node.UpdateHeight()
		right.UpdateHeight()
	}
}

func rightRotation[T Number](node *AVLNode[T]) {
	if node != nil && node.Left != nil {
		left := node.Left
		node.Left = left.Right
		left.Right = node
		node.UpdateHeight()
		left.UpdateHeight()
	}
}

func avlInsertImpl[T Number](node **AVLNode[T], data T) {
	if (*node) == nil {
		*node = NewAVLNode(data)
		return
	}

	if data < (*node).Data {
		avlInsertImpl(&(*node).Left, data)
		(*node).UpdateHeight()
		if (*node).GetBalanceFactor() == 2 {
			if (*node).Left.GetBalanceFactor() == 1 { // LL
				rightRotation(*node)
			} else { //(*node).Left.GetBalanceFactor() == -1 // LR
				leftRotation((*node).Left)
				rightRotation(*node)
			}
		}
	} else {
		avlInsertImpl(&(*node).Right, data)
		(*node).UpdateHeight()
		if (*node).GetBalanceFactor() == -2 {
			if (*node).Right.GetBalanceFactor() == -1 { // RR
				leftRotation(*node)
			} else { //(*node).Left.GetBalanceFactor() == 1 // RL
				rightRotation((*node).Right)
				leftRotation(*node)
			}
		}
	}
}

func (avl *AVL[T]) Insert(data T) {
	avlInsertImpl(&avl.Root, data)
}

// heap

type Heap[T Number] struct {
	Nodes     []T
	IsMaxHeap bool
}

func (h *Heap[T]) downAdjust(startIndex int) {
	i, j := startIndex, startIndex*2+1 // as the start index is from 0, not 1, so j is the left child of i
	for j < len(h.Nodes) {
		if h.IsMaxHeap {
			if j+1 < len(h.Nodes) && h.Nodes[j+1] > h.Nodes[j] {
				// if right child exist and the value of right child is greater than left child
				// set the greater index to right child index
				j = j + 1
			}

			if h.Nodes[i] < h.Nodes[j] {
				// if value in parent node is lower than the greater child
				// swap the two
				h.Nodes[j], h.Nodes[i] = h.Nodes[i], h.Nodes[j]
				i = j
				j = i*2 + 1
			} else {
				// take the i as the root, it is already a heap, down adjust finish
				break
			}
		} else {
			if j+1 < len(h.Nodes) && h.Nodes[j+1] < h.Nodes[j] {
				j = j + 1
			}

			if h.Nodes[j] < h.Nodes[i] {
				h.Nodes[j], h.Nodes[i] = h.Nodes[i], h.Nodes[j]
				i = j
				j = i*2 + 1
			} else {
				break
			}
		}

	}
}

// NewHeap heapify a list of values, the time complex will be O(n), not O(nlogn)
func NewHeap[T Number](values []T, isMaxHeap bool) *Heap[T] {
	h := &Heap[T]{
		Nodes:     values,
		IsMaxHeap: isMaxHeap,
	}

	for i := (len(values) - 1) / 2; i >= 0; i-- {
		// start from the last node that has child, down adjust
		h.downAdjust(i)
	}

	return h
}

func (h *Heap[T]) RemoveTop() T {
	length := len(h.Nodes)
	top := h.Nodes[0]
	h.Nodes[0] = h.Nodes[length-1] // move the last element to the front
	h.Nodes = h.Nodes[0 : length-1]
	h.downAdjust(0) // down adjust the last element (which moved to the front)
	return top
}

func (h *Heap[T]) RemoveValueOnce(v T) {
	length := len(h.Nodes)
	removeIndex := 0
	for i, node := range h.Nodes {
		if node == v {
			removeIndex = i
			break
		}
	}
	if removeIndex == 0 {
		_ = h.RemoveTop()
		return
	}
	h.Nodes[removeIndex] = h.Nodes[length-1] // move the last element to the remove index
	h.Nodes = h.Nodes[0 : length-1]
	fatherIndex := (removeIndex - 1) / 2
	if h.IsMaxHeap {
		if h.Nodes[removeIndex] > h.Nodes[fatherIndex] {
			h.upAdjust(removeIndex)
		} else {
			h.downAdjust(removeIndex)
		}
	} else {
		if h.Nodes[removeIndex] < h.Nodes[fatherIndex] {
			h.upAdjust(removeIndex)
		} else {
			h.downAdjust(removeIndex)
		}
	}
}

func (h *Heap[T]) upAdjust(startIndex int) {
	i, j := startIndex, (startIndex-1)/2 // as the index start from 0, j is the parent of i
	for j >= 0 {
		if h.IsMaxHeap {
			if h.Nodes[i] > h.Nodes[j] {
				h.Nodes[i], h.Nodes[j] = h.Nodes[j], h.Nodes[i]
				i = j
				j = (i - 1) / 2
			} else {
				break
			}
		} else {
			if h.Nodes[i] < h.Nodes[j] {
				h.Nodes[i], h.Nodes[j] = h.Nodes[j], h.Nodes[i]
				i = j
				j = (j - 1) / 2
			} else {
				break
			}
		}
	}
}

func (h *Heap[T]) Insert(data T) {
	h.Nodes = append(h.Nodes, data)
	h.upAdjust(len(h.Nodes) - 1)
}

// huffman tree

type HuffmanTree[T Number] struct {
	Root *Node[T]
}

func NewHuffmanTree[T Number](values []T) *HuffmanTree[T] {
	if len(values) == 0 {
		return nil
	}

	if len(values) == 1 {
		return &HuffmanTree[T]{
			Root: NewNode(values[0]),
		}
	}

	minHeap := NewHeap(values, false)
	rootMap := make(map[T][]*Node[T])
	var root *Node[T]
	for len(minHeap.Nodes) > 1 {
		top := minHeap.RemoveTop()
		nextTop := minHeap.RemoveTop()
		topNodeList := rootMap[top]
		var topNode *Node[T]
		if len(topNodeList) != 0 {
			topNode = topNodeList[len(topNodeList)-1]
			topNodeList = topNodeList[0 : len(topNodeList)-1]
			rootMap[top] = topNodeList
		} else {
			topNode = NewNode(top)
		}

		nextTopNodeList := rootMap[nextTop]
		var nextTopNode *Node[T]
		if len(nextTopNodeList) != 0 {
			nextTopNode = nextTopNodeList[len(nextTopNodeList)-1]
			nextTopNodeList = nextTopNodeList[0 : len(nextTopNodeList)-1]
			rootMap[nextTop] = nextTopNodeList
		} else {
			nextTopNode = NewNode(nextTop)
		}

		sum := top + nextTop
		root = NewNode(sum)
		root.Left = topNode
		root.Right = nextTopNode

		newRootList := rootMap[sum]
		if newRootList == nil {
			newRootList = make([]*Node[T], 0, 4)
		}
		newRootList = append(newRootList, root)
		rootMap[sum] = newRootList

		minHeap.Insert(sum)
	}

	return &HuffmanTree[T]{
		Root: root,
	}
}
