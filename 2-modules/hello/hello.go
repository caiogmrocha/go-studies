package main

import (
	"fmt"

	"example.com/bst"
)

func printNodeValue(n int) (error) {
	fmt.Println(n)

	return nil
}

func main() {
	var tree *bst.BST

	bst.Insert(&tree, 10)
	bst.Insert(&tree, 5)
	bst.Insert(&tree, 15)
	bst.Insert(&tree, 1)
	bst.Insert(&tree, 7)

	bst.InOrderTraversal(tree, printNodeValue)
}
