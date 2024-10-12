package bst

import (
	"errors"
	"fmt"
)

type BST struct {
	value int
	left, right *BST
}

func Insert(tree **BST, value int) (error) {
	if (value < 0) {
		return errors.New("value less than 0")
	}

	if *tree == nil {
		*tree = &BST{
			value: value,
			left: nil,
			right: nil,
		}

		return nil
	} else if (*tree).value < value {
		return Insert(&(*tree).left, value)
	} else {
		return Insert(&(*tree).right, value)
	}
}

type TraversalCallback func (int) (error)

func InOrderTraversal(tree *BST, cb TraversalCallback) {
	if (tree != nil) {
		InOrderTraversal(tree.right, cb)

		error := cb(tree.value)
		if (error != nil) {
			fmt.Println("There was an error while executing InOrderTraversak callback")
			return
		}

		InOrderTraversal(tree.left, cb)
	}
}
