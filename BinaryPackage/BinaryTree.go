package binarypackage

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (b *BinaryTree) Insert(value int) {
	if b.Root == nil {
		b.Root = createNode(value)
	} else {
		insertNode(&b.Root, value)
	}
}

func (b *BinaryTree) Search(value int) bool {
	return searchNode(&b.Root, value)
}

func (b *BinaryTree) InorderTraversal() string {
	output := ""
	traverseInorder(b.Root, &output)
	return output
}

func (b *BinaryTree) PreorderTraversal() string {
	output := ""
	traversePreorder(b.Root, &output)
	return output
}

func (b *BinaryTree) PostorderTraversal() string {
	output := ""
	traversePostorder(b.Root, &output)
	return output
}

func (b *BinaryTree) Height() int {
	return getHeight(b.Root)
}

func createNode(value int) *Node {
	var newNode Node
	newNode.Data = value
	newNode.Left = nil
	newNode.Right = nil

	return &newNode
}

func insertNode(root **Node, value int) {
	if *root == nil {
		//start at root param
		*root = createNode(value)
	} else {
		//check if value is greater than root or not
		if value > (*root).Data {
			insertNode(&(*root).Right, value)
		} else if value < (*root).Data {
			insertNode(&(*root).Left, value)
		}
	}
}

func searchNode(currentNode **Node, value int) bool {
	if *currentNode == nil {
		return false
	} else if (*currentNode).Data == value {
		return true
	} else if (*currentNode).Data > value {
		return searchNode(&(*currentNode).Left, value)
	} else if (*currentNode).Data < value {
		return searchNode(&(*currentNode).Right, value)
	}

	return false
}

func traverseInorder(currentNode *Node, output *string) {
	if currentNode != nil {
		traverseInorder(currentNode.Left, output)
		*output += fmt.Sprintf("%v ", currentNode.Data)
		traverseInorder(currentNode.Right, output)
	}
}

func traversePreorder(currentNode *Node, output *string) {
	if currentNode != nil {
		*output += fmt.Sprintf("%v ", currentNode.Data)
		traversePreorder(currentNode.Left, output)
		traversePreorder(currentNode.Right, output)
	}
}

func traversePostorder(currentNode *Node, output *string) {
	if currentNode != nil {
		traversePostorder(currentNode.Left, output)
		traversePostorder(currentNode.Right, output)
		*output += fmt.Sprintf("%v ", currentNode.Data)
	}
}

func getHeight(currentNode *Node) int {
	if currentNode == nil {
		return -1
	}
	leftHeight := getHeight(currentNode.Left)
	rightHeight := getHeight(currentNode.Right)

	return max(leftHeight, rightHeight) + 1
}

func max(A int, B int) int {
	if A > B {
		return A
	} else {
		return B
	}
}
