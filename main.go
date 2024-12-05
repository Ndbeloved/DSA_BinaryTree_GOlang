package main

import (
	"fmt"

	binarypackage "github.com/Ndbeloved/BinaryTree/BinaryPackage"
)

func main() {
	Data := []int{23, 13, 45, 19, 2, 36, 49, 22}
	find := 60
	Tree := binarypackage.NewBinaryTree()
	for _, value := range Data {
		Tree.Insert(value)
	}
	var inTree bool = Tree.Search(find)
	fmt.Printf("The value %v is in tree? %v \n", find, inTree)
	fmt.Printf("Inorder Traversal: %v \n", Tree.InorderTraversal())
	fmt.Printf("Preorder Traversal: %v \n", Tree.PreorderTraversal())
	fmt.Printf("Postorder Traversal: %v \n", Tree.PostorderTraversal())
	fmt.Printf("length of Tree: %v \n", Tree.Height())
}
