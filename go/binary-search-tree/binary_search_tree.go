package binarysearchtree

type BinarySearchTree struct {
	left  *BinarySearchTree
	data  int
	right *BinarySearchTree
}

func NewBst(i int) *BinarySearchTree {
	return &BinarySearchTree{data: i}
}

func (bst *BinarySearchTree) Insert(i int) *BinarySearchTree {
	if bst == nil {
		return &BinarySearchTree{data: i}
	}

	if i <= bst.data {
		bst.left = bst.left.Insert(i)
	} else {
		bst.right = bst.right.Insert(i)
	}

	return bst
}

func (bst *BinarySearchTree) SortedData() []int {
	if bst == nil {
		return []int{}
	}

	var result []int

	result = append(result, bst.left.SortedData()...)
	result = append(result, bst.data)
	result = append(result, bst.right.SortedData()...)

	return result
}
