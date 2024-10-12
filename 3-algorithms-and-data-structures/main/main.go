package main

import (
	"ads/bst"
	"errors"
	"fmt"
)

func printNodeValue(n *bst.Node) (error) {
	if (n == nil) {
		return errors.New("n is nil")
	}

	fmt.Println(n.Value)

	return nil
}

func main() {
		tree := &bst.BST{}

		tree.Insert(10)
		tree.Insert(5)
		tree.Insert(15)
		tree.Insert(1)
		tree.Insert(7)

		// tree.InOrderTraversal(printNodeValue)

		tree.Remove(5)

		tree.InOrderTraversal(printNodeValue)
}
