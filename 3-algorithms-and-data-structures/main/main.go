package main

import (
	"ads/graph"
	"fmt"
)

func main() {
	g := &graph.Graph[string]{}

	g.AddVertex(&graph.Vertex[string]{Data: "A"})
	g.AddVertex(&graph.Vertex[string]{Data: "B"})
	g.AddVertex(&graph.Vertex[string]{Data: "C"})
	g.AddVertex(&graph.Vertex[string]{Data: "D"})
	g.AddVertex(&graph.Vertex[string]{Data: "E"})

	g.AddAdjacence(g.Vertices[0], g.Vertices[1])
	g.AddAdjacence(g.Vertices[1], g.Vertices[2])
	g.AddAdjacence(g.Vertices[2], g.Vertices[3])
	g.AddAdjacence(g.Vertices[3], g.Vertices[4])
	g.AddAdjacence(g.Vertices[4], g.Vertices[0])

	fmt.Println("Vertices:")

	for k,v := range g.Vertices {
		fmt.Printf("%d - %s\n", k, v.Data)
	}
}

// func printNodeValue(n *bst.Node) (error) {
// 	if (n == nil) {
// 		return errors.New("n is nil")
// 	}

// 	fmt.Println(n.Value)

// 	return nil
// }

// func testBSTPackage() {
// 	tree := &bst.BST{}

// 	tree.Insert(10)
// 	tree.Insert(5)
// 	tree.Insert(15)
// 	tree.Insert(1)
// 	tree.Insert(7)

// 	tree.InOrderTraversal(printNodeValue)

// 	tree.Remove(5)

// 	tree.InOrderTraversal(printNodeValue)
// }
