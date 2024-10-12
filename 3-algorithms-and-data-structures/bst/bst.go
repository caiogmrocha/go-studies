package bst

import (
	"errors"
	"fmt"
)

type BST struct {
	root *Node
}

type Node struct {
	Value int
	Left, Right *Node
}

func (tree *BST) Insert(value int) (error) {
	if (value < 0) {
		return errors.New("value less than 0")
	}

	if tree.root == nil {
		tree.root = &Node{
			Value: value,
		}

		return nil
	} else {
		return tree.root.insert(value)
	}
}

func (node *Node) insert(value int) (error) {
	if node.Value < value {
		if node.Left == nil {
			node.Left = &Node{Value: value}

			return nil
		} else {
			return node.Left.insert(value)
		}
	} else {
		if node.Right == nil {
			node.Right = &Node{Value: value}

			return nil
		} else {
			return node.Right.insert(value)
		}
	}
}

func (node *Node) getRightmostNode() *Node {
	iterator := node

	for iterator.Right != nil {
		iterator = iterator.Right
	}

	return iterator
}

func (tree *BST) Remove(value int) (error) {
	if value < 0 {
		return errors.New("value less than 0")
	}

	if tree.root.Value == value {
		tree.root = nil

		return nil
	} else if (tree.root.Value < value) {
		if tree.root.Left != nil {
			return tree.root.remove(value)
		}

		return nil
	} else {
		if tree.root.Right != nil {
			return tree.root.remove(value)
		}

		return nil
	}
}

func (node *Node) remove(value int) (error) {
	if (node == nil) {
		return errors.New("tree is nil or value not found")
	} else if (node.Left != nil && node.Left.Value == value) {
		if (node.Left.Left == nil && node.Left.Right == nil) {
			node.Left = nil
		} else if (node.Left.Left != nil && node.Left.Right == nil) {
			node.Left = node.Left.Left
		} else if (node.Left.Left == nil && node.Left.Right != nil) {
			node.Left = node.Left.Right
		} else {
			rightmostNode := node.Left.Left.getRightmostNode()

			node.Left.Value = rightmostNode.Value

			node.Left.remove(rightmostNode.Value)
		}

		return nil
	} else if (node.Right != nil && node.Right.Value == value) {
		if (node.Right.Left == nil && node.Right.Right == nil) {
			node.Right = nil
		} else if (node.Right.Left != nil && node.Right.Right == nil) {
			node.Right = node.Right.Left
		} else if (node.Right.Left == nil && node.Right.Right != nil) {
			node.Right = node.Right.Right
		} else {
			rightmostNode := node.Right.Left.getRightmostNode()

			node.Right.Value = rightmostNode.Value

			node.Right.remove(rightmostNode.Value)
		}

		return nil
	} else if (node.Value < value) {
		return node.Left.remove(value)
	} else {
		return node.Right.remove(value)
	}
}

type TraversalCallback func (*Node) (error)


func (tree *BST) PreOrderTraversal(cb TraversalCallback) {
	if (tree != nil) {
		tree.root.Right.PreOrderTraversal(cb)
		tree.root.Left.PreOrderTraversal(cb)

		error := cb(tree.root)
		if (error != nil) {
			fmt.Println("There was an error while executing PreOrderTraversal callback")
			return
		}
	}
}
func (node *Node) PreOrderTraversal(cb TraversalCallback) {
	if (node != nil) {
		node.Right.PreOrderTraversal(cb)
		node.Left.PreOrderTraversal(cb)

		error := cb(node)
		if (error != nil) {
			fmt.Println("There was an error while executing PreOrderTraversal callback")
			return
		}
	}
}

func (tree *BST) InOrderTraversal(cb TraversalCallback) {
	if (tree != nil) {
		tree.root.Right.inOrderTraversal(cb)

		error := cb(tree.root)
		if (error != nil) {
			fmt.Println("There was an error while executing InOrderTraversak callback")
			return
		}

		tree.root.Left.inOrderTraversal(cb)
	}
}

func (node *Node) inOrderTraversal(cb TraversalCallback) {
	if (node != nil) {
		node.Right.inOrderTraversal(cb)

		error := cb(node)
		if (error != nil) {
			fmt.Println("There was an error while executing InOrderTraversak callback")
			return
		}

		node.Left.inOrderTraversal(cb)
	}
}

func (tree *BST) PostOrderTraversal(cb TraversalCallback) {
	if (tree != nil) {
		error := cb(tree.root)
		if (error != nil) {
			fmt.Println("There was an error while executing PreOrderTraversal callback")
			return
		}

		tree.root.Right.postOrderTraversal(cb)
		tree.root.Left.postOrderTraversal(cb)
	}
}

func (node *Node) postOrderTraversal(cb TraversalCallback) {
	if (node != nil) {
		error := cb(node)
		if (error != nil) {
			fmt.Println("There was an error while executing PreOrderTraversal callback")
			return
		}

		node.Right.postOrderTraversal(cb)
		node.Left.postOrderTraversal(cb)
	}
}
