package bst

import (
	"errors"
	"fmt"
)

type BST struct {
	Value int
	Left, Right *BST
}

func Insert(tree **BST, value int) (error) {
	if (value < 0) {
		return errors.New("value less than 0")
	}

	if *tree == nil {
		*tree = &BST{
			Value: value,
			Left: nil,
			Right: nil,
		}

		return nil
	} else if (*tree).Value < value {
		return Insert(&(*tree).Left, value)
	} else {
		return Insert(&(*tree).Right, value)
	}
}

func getRightmostNode(tree **BST) *BST {
	iterator := *tree

	for iterator.Right != nil {
		iterator = iterator.Right
	}

	return iterator
}

func Remove(tree **BST, value int) (error) {
	if (*tree == nil) {
		return errors.New("tree is nil or value not found")
	} else if ((*tree).Value == value) {
		if ((*tree).Left == nil && (*tree).Right == nil) {
			*tree = nil
		} else if ((*tree).Left != nil && (*tree).Right == nil) {
			*tree = (*tree).Left
		} else if ((*tree).Left == nil && (*tree).Right != nil) {
			*tree = (*tree).Right
		} else {
			rightmostNode := getRightmostNode(&(*tree).Left)

			(*tree).Value = rightmostNode.Value

			Remove(&(*tree).Left, rightmostNode.Value)
		}

		return nil
	} else if ((*tree).Value < value) {
		return Remove(&(*tree).Left, value)
	} else {
		return Remove(&(*tree).Right, value)
	}
}

type TraversalCallback func (*BST) (error)

func PreOrderTraversal(tree *BST, cb TraversalCallback) {
	if (tree != nil) {
		PreOrderTraversal(tree.Right, cb)
		PreOrderTraversal(tree.Left, cb)

		error := cb(tree)
		if (error != nil) {
			fmt.Println("There was an error while executing PreOrderTraversal callback")
			return
		}
	}
}

func InOrderTraversal(tree *BST, cb TraversalCallback) {
	if (tree != nil) {
		InOrderTraversal(tree.Right, cb)

		error := cb(tree)
		if (error != nil) {
			fmt.Println("There was an error while executing InOrderTraversak callback")
			return
		}

		InOrderTraversal(tree.Left, cb)
	}
}

func PostOrderTraversal(tree *BST, cb TraversalCallback) {
	if (tree != nil) {
		error := cb(tree)
		if (error != nil) {
			fmt.Println("There was an error while executing PreOrderTraversal callback")
			return
		}

		PostOrderTraversal(tree.Right, cb)
		PostOrderTraversal(tree.Left, cb)
	}
}
