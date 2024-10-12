package main

import (
	"errors"
	"fmt"

	"example.com/bst"
)

func printNodeValue(n *bst.BST) (error) {
	if (n == nil) {
		return errors.New("n is nil")
	}

	fmt.Println(n.Value)

	return nil
}

func main() {
	var tree *bst.BST

	bst.Insert(&tree, 10)
	bst.Insert(&tree, 5)
	bst.Insert(&tree, 15)
	bst.Insert(&tree, 1)
	bst.Insert(&tree, 7)

	bst.PreOrderTraversal(tree, printNodeValue)
	fmt.Println()
	bst.InOrderTraversal(tree, printNodeValue)
	fmt.Println()
	bst.PostOrderTraversal(tree, printNodeValue)
	fmt.Println()
}
