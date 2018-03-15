package binarytree

type treeNode struct {
	TElement   interface{}
	parent     *treeNode
	leftChild  *treeNode
	rightChild *treeNode
	level      int
}

// BinaryTree creates a Binary Tree
type BinaryTree struct {
	treeNode *treeNode
}

// NewBinaryTree instantiates new tree
func NewBinaryTree(root interface{}) *BinaryTree {
	tree := new(BinaryTree)

	return tree
}
