package binarytree

import "github.com/thomasCM23/learnGoLang/Queue"

type treeNode struct {
	TElement   interface{}
	parent     *treeNode
	leftChild  *treeNode
	rightChild *treeNode
	level      int
}

// BinaryTree creates a Binary Tree
type BinaryTree struct {
	treeNode       *treeNode
	addToLeftNode  *queue.Queue
	addToRightNode *queue.Queue
}

// GetParentNode return the parent of the past child
func (b *BinaryTree) getParentNode(child *treeNode) *treeNode {
	return child.parent
}

// AppendLeftNode adds node to left child of passed parent node
func (b *BinaryTree) appendLeftNode(parent *treeNode, item interface{}) *treeNode {
	leftChildNode := new(treeNode)
	leftChildNode.parent = parent
	leftChildNode.TElement = item
	leftChildNode.level = parent.level + 1
	parent.leftChild = leftChildNode
	return leftChildNode
}

// GetLeftChild for parent node passed
func (b *BinaryTree) getLeftChild(parent *treeNode) *treeNode {
	return parent.leftChild
}

// AppendRightNode adds node to left child of passed parent node
func (b *BinaryTree) appendRightNode(parent *treeNode, item interface{}) *treeNode {
	rightChildNode := new(treeNode)
	rightChildNode.parent = parent
	rightChildNode.TElement = item
	rightChildNode.level = parent.level + 1
	parent.rightChild = rightChildNode
	return rightChildNode
}

// GetRightChild for parent node passed
func (b *BinaryTree) getRightChild(parent *treeNode) *treeNode {
	return parent.rightChild
}

// Append item to tree to form a balance binary tree
func (b *BinaryTree) Append(item interface{}) {
	if b.addToRightNode.Size() > 1 {
		parentNode := (b.addToRightNode.Dequeue()).(*treeNode)
		rightNode := b.appendLeftNode(parentNode, item)
		b.addToLeftNode.Enqueue(rightNode)
	} else {
		parentNode := (b.addToLeftNode.Dequeue()).(*treeNode)
		leftNode := b.appendLeftNode(parentNode, item)
		b.addToRightNode.Enqueue(parentNode)
		b.addToLeftNode.Enqueue(leftNode)
	}
}

// NewBinaryTree instantiates new tree
func NewBinaryTree(root interface{}) (tree *BinaryTree) {
	tree = new(BinaryTree)
	tree.treeNode = new(treeNode)
	tree.addToLeftNode = queue.NewQueue(tree.treeNode)
	tree.addToRightNode = queue.NewQueue()
	tree.treeNode.level = 0
	tree.treeNode.parent = nil
	tree.treeNode.TElement = root
	return tree
}
